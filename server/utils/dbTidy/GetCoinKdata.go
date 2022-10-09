package dbTidy

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCoinKdata() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("CoinTicker")
	defer db.Close()

	FK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, FK)
	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Ticker dbType.CoinTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)

		FetchKdata(Ticker.TickerVol)

	}

	if err != nil {
		resErr := fmt.Errorf("GetCoinKdata 失败 %+v", err)
		global.LogErr(resErr)
		return
	}
}

func FetchKdata(List []mOKX.TypeTicker) {
	for _, val := range List {
		fmt.Println(val.InstID)
	}
}
