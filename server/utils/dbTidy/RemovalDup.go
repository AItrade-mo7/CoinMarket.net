package dbTidy

import (
	"fmt"
	"os"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimeUnixType struct {
	TimeUnix int64  `bson:"TimeUnix"`
	TimeStr  string `bson:"TimeStr"`
	TimeID   string `bson:"TimeID"`
}

func RemovalDup() {
	TimeUnixArr_file := mStr.Join(config.Dir.JsonData, "/TimeUnixArr", ".json")
	var TimeUnixArr []TimeUnixType

	fileCont, _ := os.ReadFile(TimeUnixArr_file)
	jsoniter.Unmarshal(fileCont, &TimeUnixArr)

	if len(TimeUnixArr) > 2000 {
		CheckRepeat(TimeUnixArr)
		return
	}

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

	FK := bson.D{}

	cursor, err := db.Table.Find(db.Ctx, FK, findOpt)
	fmt.Println(cursor, err)

	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Ticker dbType.CoinTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)

		TimeObj := TimeUnixType{
			TimeUnix: Ticker.TimeUnix,
			TimeStr:  Ticker.TimeStr,
			TimeID:   Ticker.TimeID,
		}

		TimeUnixArr = append(TimeUnixArr, TimeObj)

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

		global.Run.Println("==结束==", Ticker.TimeStr, len(Ticker.Kdata), len(BtcList), timeC, errArr)
	}

	if err != nil {
		global.LogErr(err)
	}

	mFile.Write(TimeUnixArr_file, mJson.ToStr(TimeUnixArr))
}

func CheckRepeat(list []TimeUnixType) {
	fmt.Println("开始检查重复", len(list))
	RepeatTimeID_file := mStr.Join(config.Dir.JsonData, "/RepeatTimeID", ".json")
	RepeatIndex_file := mStr.Join(config.Dir.JsonData, "/RepeatIndex", ".json")

	timeMap := make(map[string]TimeUnixType)
	var RepeatTimeID []string
	var RepeatIndex []int

	fileCont, _ := os.ReadFile(RepeatTimeID_file)
	jsoniter.Unmarshal(fileCont, &RepeatTimeID)

	if len(RepeatTimeID) > 0 {
		RemoveRepeat(RepeatTimeID)
		return
	}

	for key, val := range list {
		TimeID := val.TimeID
		_, ok := timeMap[TimeID]
		if ok {
			RepeatTimeID = append(RepeatTimeID, val.TimeID)
			RepeatIndex = append(RepeatIndex, key)
		} else {
			timeMap[TimeID] = val
		}
	}

	mFile.Write(RepeatTimeID_file, mJson.ToStr(RepeatTimeID))
	mFile.Write(RepeatIndex_file, mJson.ToStr(RepeatIndex))
}

func RemoveRepeat(timeIDList []string) {
	fmt.Println("开始删除重复数据", len(timeIDList))
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
	for _, val := range timeIDList {
		findFK := bson.D{{
			Key:   "TimeID",
			Value: val,
		}}
		cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)

		fmt.Println("开始处理", cursor, err)

		IDArr := []string{}

		for cursor.Next(db.Ctx) {
			var curData map[string]any
			cursor.Decode(&curData)
			var Ticker dbType.CoinTickerTable
			jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)

			IDArr = append(IDArr, Ticker.TimeID)

			BtcList := Ticker.Kdata["BTC-USDT"]
			global.Run.Println("==结束==", curData["_id"], Ticker.TimeStr, Ticker.TimeID, len(Ticker.Kdata), len(BtcList))
		}

		if len(IDArr) > 1 {
			fmt.Println("开始删除", IDArr)
			findOneFK := bson.D{{
				Key:   "TimeID",
				Value: IDArr[0],
			}}
			db.Table.DeleteOne(db.Ctx, findOneFK)
		}

		cursor2, err := db.Table.Find(db.Ctx, findFK, findOpt)

		fmt.Println("处理完毕", cursor2, err)

		if err != nil {
			global.LogErr(err)
		}

	}
}
