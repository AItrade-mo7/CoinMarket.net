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
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimeUnixType struct {
	TimeUnix int64  `bson:"TimeUnix"`
	TimeStr  string `bson:"TimeStr"`
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
	timeMap := make(map[string]TimeUnixType)
	RepeatArr := []TimeUnixType{}
	RepeatIndex := []int{}

	for key, val := range list {
		TimeUnix := val.TimeUnix
		T := mTime.MsToTime(TimeUnix, "0")
		timeStr := T.Format("2006-01-02T15:04")
		_, ok := timeMap[timeStr]
		if ok {
			RepeatArr = append(RepeatArr, val)
			RepeatIndex = append(RepeatIndex, key)
		} else {
			timeMap[timeStr] = val
		}

	}

	RepeatArr_file := mStr.Join(config.Dir.JsonData, "/RepeatArr", ".json")
	RepeatIndex_file := mStr.Join(config.Dir.JsonData, "/RepeatIndex", ".json")

	mFile.Write(RepeatArr_file, mJson.ToStr(RepeatArr))
	mFile.Write(RepeatIndex_file, mJson.ToStr(RepeatIndex))
}
