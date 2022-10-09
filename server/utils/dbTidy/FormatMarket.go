package dbTidy

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMarketTickerData() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()

	dbCoin := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("CoinTicker")
	defer dbCoin.Close()

	FK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, FK)

	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Ticker dbType.MarketTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)

		InsertCoinTicker(dbCoin, Ticker)

	}

	if err != nil {
		fmt.Println(err)
	}
}

func InsertCoinTicker(db *mMongo.DB, Ticker dbType.MarketTickerTable) {
	var CoinTickerData dbType.CoinTickerTable
	CoinTickerData.TickerVol = Ticker.List
	CoinTickerData.TimeUnix = Ticker.List[0].Ts
	CoinTickerData.TimeStr = mTime.UnixFormat(mStr.ToStr(Ticker.List[0].Ts))

	FK := bson.D{{
		Key:   "TimeUnix",
		Value: CoinTickerData.TimeUnix,
	}}
	UK := bson.D{}
	mStruct.Traverse(CoinTickerData, func(key string, val any) {
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

	var result dbType.CoinTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&result)

	var err error
	lType := ""
	if result.TimeUnix > 0 {
		lType = "更新"
		global.Run.Println("进行数据更新", CoinTickerData.TimeStr)
		_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
	} else {
		lType = "插入"
		global.Run.Println("进行数据插入", CoinTickerData.TimeStr)
		_, err = db.Table.InsertOne(db.Ctx, CoinTickerData)
	}
	if err != nil {
		resErr := fmt.Errorf(lType+"数据失败 %+v", err)
		global.LogErr(resErr)
		return
	}
}
