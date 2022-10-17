package dbTask

import (
	"fmt"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
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
		global.Run.Println("==结束==", Ticker.TimeID, Ticker.TimeStr, len(Ticker.Kdata), len(BtcList), timeC, errArr)
	}

	if err != nil {
		global.LogErr(err)
	}
}

func FetchKdata(dbTicker dbType.CoinTickerTable) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)

	for _, val := range dbTicker.TickerVol {
		kdata_list := dbTicker.Kdata[val.InstID]

		fmt.Println("db", len(kdata_list))
		if len(kdata_list) != 100 {
			time.Sleep(time.Second / 3)
			kdata_list = kdata.GetHistoryKdata(kdata.HistoryKdataParam{
				InstID:  val.InstID,
				Current: 0,
				Size:    100,
				After:   val.Ts,
			})
			fmt.Println("fetch", len(kdata_list))
		}
		KdataList[val.InstID] = kdata_list
		global.Run.Println("填充结束", val.InstID, len(KdataList[val.InstID]), kdata_list[0].TimeStr, kdata_list[len(kdata_list)-1].TimeStr)
	}
	return KdataList
}
