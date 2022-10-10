package dbTidy

import (
	"fmt"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FormatMarket() {
	fmt.Println("开始执行")
	inst.Start()

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  3241 * 10000 * 60,
	}).Connect().Collection("CoinTicker")
	defer db.Close()

	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	cursor, err := db.Table.Find(db.Ctx, bson.D{}, findOpt)

	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Ticker dbType.CoinTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)
		Ticker.TimeStr = mTime.UnixFormat(mStr.ToStr(Ticker.TimeUnix))

		kdata_list := FetchKdata(Ticker)

		Ticker.Kdata = kdata_list

		FK := bson.D{{
			Key:   "TimeUnix",
			Value: Ticker.TimeUnix,
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

		var newTicker dbType.CoinTickerTable
		db.Table.FindOne(db.Ctx, FK).Decode(&newTicker)

		if len(newTicker.Kdata) != len(newTicker.TickerVol) || len(newTicker.Kdata["BTC-USDT"]) < 280 {
			global.LogErr("==错误==", newTicker.TimeStr, len(newTicker.Kdata), len(newTicker.Kdata["BTC-USDT"]))
		} else {
			global.Run.Println("==结束==", newTicker.TimeStr, len(newTicker.Kdata), len(newTicker.Kdata["BTC-USDT"]))
		}
	}

	if err != nil {
		global.LogErr(err)
	}
}

func FetchKdata(dbTicker dbType.CoinTickerTable) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)

	for _, val := range dbTicker.TickerVol {
		kdata_list := dbTicker.Kdata[val.InstID]

		if len(kdata_list) < 290 {
			time.Sleep(time.Second / 3)
			kdata_list = kdata.GetHistory300List(kdata.History300Param{
				InstID: val.InstID,
				After:  val.Ts,
			})

		}

		KdataList[val.InstID] = kdata_list
		global.Run.Println("填充结束", val.InstID, len(KdataList[val.InstID]))
	}

	return KdataList
}
