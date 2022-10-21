package dbTask

import (
	"fmt"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FormatTickerObj struct {
	TickerDB *mMongo.DB
}

func NewFormat() *FormatTickerObj {
	inst.Start()

	var NewFormatObj FormatTickerObj
	Timeout := 4000 * 60
	NewFormatObj.TickerDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("CoinTicker")

	return &NewFormatObj
}

func (_this *FormatTickerObj) TickerDBTraverse() {
	db := _this.TickerDB
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	findFK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	for cursor.Next(db.Ctx) {
		var Ticker dbType.CoinTickerTable
		cursor.Decode(&Ticker)

		Ticker.Kdata = make(map[string][]mOKX.TypeKd)
		Ticker.Kdata = FetchKdata(Ticker)
		Ticker.TimeUnix = Ticker.TickerVol[0].Ts
		Ticker.TimeStr = mTime.UnixFormat(Ticker.TimeUnix)
		Ticker.TimeID = mOKX.GetTimeID(Ticker.TimeUnix)

		Ticker.TickerVol = FormatTickerVol(Ticker.TickerVol)

		FK := bson.D{{
			Key:   "TimeID",
			Value: Ticker.TimeID,
		}}
		UK := bson.D{}
		mStruct.Traverse(Ticker, func(key string, val any) {
			UK = append(UK, bson.E{
				Key: "$set",
				Value: bson.D{
					{
						Key:   key,
						Value: val,
					},
				},
			})
		})
		upOpt := options.Update()
		upOpt.SetUpsert(true)
		_, err := db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
		if err != nil {
			global.LogErr("数据更插失败", err)
		}

		BtcList := Ticker.Kdata["BTC-USDT"]
		var timeC int64
		if len(BtcList) > 0 {
			BtcNow := BtcList[len(BtcList)-1]
			timeC = BtcNow.TimeUnix - Ticker.TimeUnix
		}
		errArr := []any{}
		for key, val := range Ticker.Kdata {
			if len(val) != config.KdataLen {
				errArr = append(errArr, key)
				errArr = append(errArr, len(val))
			}
		}
		global.Run.Println("==结束==", Ticker.TimeID, Ticker.TimeStr, len(Ticker.TickerVol), len(Ticker.Kdata), len(BtcList), timeC, errArr)
	}

	if err != nil {
		global.LogErr(err)
	}
}

func FetchKdata(dbTicker dbType.CoinTickerTable) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)

	for _, val := range dbTicker.TickerVol {
		kdata_list := dbTicker.Kdata[val.InstID]

		if len(kdata_list) != 100 {
			time.Sleep(time.Second / 3)
			kdata_list = kdata.GetHistoryKdata(kdata.HistoryKdataParam{
				InstID:  val.InstID,
				Current: 0,
				Size:    100,
				After:   val.Ts,
			})
		}
		KdataList[val.InstID] = kdata_list
		global.Run.Println("填充结束", val.InstID, len(KdataList[val.InstID]), kdata_list[0].TimeStr, kdata_list[len(kdata_list)-1].TimeStr)
	}
	return KdataList
}

func FormatTickerVol(TickerVol []mOKX.TypeTicker) []mOKX.TypeTicker {
	NewTickerVol := []mOKX.TypeTicker{}

	for _, Ticker := range TickerVol {
		NewTicker := Ticker
		NewTicker.TimeUnix = mTime.ToUnixMsec(mTime.MsToTime(Ticker.Ts, "0"))
		NewTicker.TimeStr = mTime.UnixFormat(Ticker.TimeUnix)
		Ticker.SWAP = mOKX.TypeInst{}
		Ticker.SPOT = mOKX.TypeInst{}
		if len(Ticker.InstID) > 3 {
			for _, SWAP := range okxInfo.SWAP_inst {
				if SWAP.Uly == Ticker.InstID {
					Ticker.SWAP = SWAP
					break
				}
			}
			for _, SPOT := range okxInfo.SPOT_inst {
				if SPOT.InstID == Ticker.InstID {
					Ticker.SPOT = SPOT
					break
				}
			}
		}
		// 上架小于36天的不计入榜单
		diffOnLine := mCount.Sub(mStr.ToStr(Ticker.Ts), Ticker.SPOT.ListTime)

		fmt.Println()
		if mCount.Le(diffOnLine, "32") > 0 {
			NewTickerVol = append(NewTickerVol, NewTicker)
			global.Run.Println("榜单填充结束", NewTicker.InstID)
		} else {
			global.Run.Println("上架时间太短-过滤", diffOnLine)
		}
	}

	return NewTickerVol
}
