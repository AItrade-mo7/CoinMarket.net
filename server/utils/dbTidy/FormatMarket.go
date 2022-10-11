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
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FormatMarket() {
	// global.Email(global.EmailOpt{
	// 	To: []string{
	// 		"meichangliang@mo7.cc",
	// 	},
	// 	Subject:  "ServeStart",
	// 	Template: tmpl.SysEmail,
	// 	SendData: tmpl.SysParam{
	// 		Message: "开始执行脚本",
	// 		SysTime: time.Now(),
	// 	},
	// }).Send()
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
		"TimeUnix": -1,
	})

	findFK := bson.D{{
		Key:   "TimeUnix",
		Value: 1662453885809,
	}}

	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)

	fmt.Println(cursor)

	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Ticker dbType.CoinTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)
		T := mTime.MsToTime(Ticker.TimeUnix, "0")
		timeStr := T.Format("2006-01-02T15:04")
		Ticker.TimeID = timeStr

		FK := bson.D{{
			Key:   "_id",
			Value: curData["_id"],
		}}
		UK := bson.D{}

		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "TimeID",
					Value: timeStr,
				},
			},
		})

		upOpt := options.Update()
		// upOpt.SetUpsert(true)
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
			if len(val) != 300 {
				errArr = append(errArr, key)
				errArr = append(errArr, len(val))
			}
		}

		global.Run.Println("==结束==", Ticker.TimeStr, Ticker.TimeID, len(Ticker.Kdata), len(BtcList), timeC, errArr)
	}

	if err != nil {
		global.LogErr(err)
	}

	// global.Email(global.EmailOpt{
	// 	To: []string{
	// 		"meichangliang@mo7.cc",
	// 	},
	// 	Subject:  "ServeStart",
	// 	Template: tmpl.SysEmail,
	// 	SendData: tmpl.SysParam{
	// 		Message: "脚本执行结束",
	// 		SysTime: time.Now(),
	// 	},
	// }).Send()
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
