package dbTidy

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckCoinKdata() {
	fmt.Println("开始检查重复数据")
	CoinName := "BTC"
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  99999,
	}).Connect().Collection(CoinName + "USDT")
	defer global.Run.Println("关闭数据库连接" + CoinName)
	defer db.Close()

	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})
	FK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, FK, findOpt)

	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Kdata mOKX.TypeKd
		jsoniter.Unmarshal(mJson.ToJson(curData), &Kdata)

		global.Run.Println("==结束==", Kdata.InstID, Kdata.TimeStr, Kdata.C)
	}

	if err != nil {
		global.LogErr(err)
	}
}
