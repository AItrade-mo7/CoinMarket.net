package dbTask

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi"
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
	okxApi.SetInst() // 获取并设置交易产品信息

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
		"TimeUnix": -1,
	})

	findFK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	for cursor.Next(db.Ctx) {
		var CoinTicker dbType.CoinTickerTable
		cursor.Decode(&CoinTicker)

		CoinTicker.TickerVol = FormatTickerVol(CoinTicker.TickerVol)
		CoinTicker.Kdata = FetchKdata(CoinTicker)

		if len(CoinTicker.TickerVol) < 1 {
			global.Run.Println("没有榜单数据", CoinTicker.TimeStr, len(CoinTicker.TickerVol), len(CoinTicker.Kdata))
			continue
		}

		CoinTicker.TimeUnix = CoinTicker.TickerVol[0].TimeUnix
		CoinTicker.TimeStr = mTime.UnixFormat(CoinTicker.TimeUnix)
		CoinTicker.TimeID = mOKX.GetTimeID(CoinTicker.TimeUnix)

		FK := bson.D{{
			Key:   "TimeID",
			Value: CoinTicker.TimeID,
		}}
		UK := bson.D{}
		mStruct.Traverse(CoinTicker, func(key string, val any) {
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

		BtcList := CoinTicker.Kdata["BTC-USDT"]
		var timeC int64
		if len(BtcList) > 0 {
			BtcNow := BtcList[len(BtcList)-1]
			timeC = BtcNow.TimeUnix - CoinTicker.TimeUnix
		}
		errArr := []any{}
		for key, val := range CoinTicker.Kdata {
			if len(val) != config.KdataLen {
				errArr = append(errArr, key)
				errArr = append(errArr, len(val))
			}
		}
		global.Run.Println("==结束==", CoinTicker.TimeID, CoinTicker.TimeStr, len(CoinTicker.TickerVol), len(CoinTicker.Kdata), len(BtcList), timeC, errArr)
	}

	if err != nil {
		global.LogErr(err)
	}
}

func FetchKdata(dbTicker dbType.CoinTickerTable) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)

	for _, val := range dbTicker.TickerVol {
		kdata_list := dbTicker.Kdata[val.InstID]
		if len(kdata_list) < 50 {
			kdata_list = mOKX.GetKdata(mOKX.GetKdataOpt{
				InstID: val.InstID,
				Page:   0,
				After:  val.TimeUnix,
			})
		}
		time.Sleep(time.Second / 3)
		KdataList[val.InstID] = kdata_list
		global.Run.Println("填充结束", val.InstID, len(KdataList[val.InstID]), kdata_list[0].TimeStr, kdata_list[len(kdata_list)-1].TimeStr)
	}
	return KdataList
}

func FormatTickerVol(TickerVol []mOKX.TypeTicker) []mOKX.TypeTicker {
	NewTickerVol := []mOKX.TypeTicker{}

	for _, Ticker := range TickerVol {
		NewTicker := Ticker

		NewTicker.TimeUnix = Ticker.TimeUnix
		NewTicker.TimeStr = mTime.UnixFormat(NewTicker.TimeUnix)

		if len(NewTicker.OkxVolRose) < 1 || len(NewTicker.BinanceVolRose) < 1 {
			NewTicker.OkxVolRose = mCount.PerCent(NewTicker.OKXVol24H, NewTicker.Volume)
			NewTicker.BinanceVolRose = mCount.PerCent(NewTicker.BinanceVol24H, NewTicker.Volume)
		}
		// 标记
		if len(Ticker.InstID) > 3 {
			NewTicker.SWAP = okxInfo.Inst[mStr.Join(Ticker.CcyName, config.SWAP_suffix)]
			NewTicker.SPOT = okxInfo.Inst[Ticker.InstID]
		}

		if len(NewTicker.SPOT.ListTime) < 4 || len(NewTicker.SWAP.ListTime) < 4 {
			EndEmail("时间错误2")
			global.Run.Println("时间错误2", NewTicker.TimeUnix)
			panic("时间太小了2")
		}

		// 上架小于 32 天的不计入榜单
		diffOnLine := mCount.Sub(mStr.ToStr(Ticker.TimeUnix), Ticker.SWAP.ListTime)
		diffDay := mCount.Div(diffOnLine, mTime.UnixTime.Day)
		if mCount.Le(diffDay, "32") < 0 {
			global.Run.Println("上架时间太短-过滤", diffOnLine)
		} else {
			NewTickerVol = append(NewTickerVol, NewTicker)
			global.Run.Println("榜单填充结束", NewTicker.InstID, NewTicker.SPOT.ListTime, NewTicker.SWAP.ListTime, NewTicker.TimeUnix)
		}
	}

	return NewTickerVol
}
