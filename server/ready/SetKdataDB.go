package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetCoinKdataDB(CoinName string) {
	if CoinName == "BTC" || CoinName == "ETH" {
	} else {
		return
	}

	tableName := CoinName + "USDT"
	InstID := CoinName + "-USDT"

	list := okxInfo.TickerKdata[InstID]

	Timeout := len(list) * 10
	if Timeout < 100 {
		Timeout = 100
	}

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection(tableName)
	defer global.Run.Println("关闭数据库连接", tableName)
	defer db.Close()

	for _, Kd := range list {
		FK := bson.D{{
			Key:   "TimeUnix",
			Value: Kd.TimeUnix,
		}}
		UK := bson.D{}
		mStruct.Traverse(Kd, func(key string, val any) {
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
			global.LogErr(tableName+"数据更插失败  %+v", err)
		}
	}
}

func SetCoinTickerDB() {
	Timeout := len(okxInfo.TickerVol) * 20

	if Timeout < 100 {
		Timeout = 100
	}
	global.Run.Println("开始存储 CoinTicker")
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("CoinTicker")
	defer global.Run.Println("关闭数据库连接 CoinTicker")
	defer db.Close()

	// 获取当前的榜单数据并拼接
	Ticker := dbType.JoinCoinTicker(tickerAnaly.TickerAnalyParam{
		TickerVol:   okxInfo.TickerVol,
		TickerKdata: okxInfo.TickerKdata,
	})

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
		resErr := fmt.Errorf("数据更插失败 CoinTicker %+v", err)
		global.LogErr(resErr)
	}
}

func SetTickerAnalyDB() {
	Timeout := len(okxInfo.TickerAnaly.AnalySingle) * 20

	if Timeout < 100 {
		Timeout = 100
	}
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("TickerAnaly")
	defer global.Run.Println("关闭数据库连接 TickerAnaly")
	defer db.Close()

	TickerAnaly := okxInfo.TickerAnaly

	FK := bson.D{{
		Key:   "TimeID",
		Value: TickerAnaly.TimeID,
	}}
	UK := bson.D{}
	mStruct.Traverse(TickerAnaly, func(key string, val any) {
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
		global.LogErr("数据更插失败 TickerAnaly %+v", err)
	}

	global.Run.Println(
		"CoinTicker更插完毕",
		TickerAnaly.TimeStr,
		len(TickerAnaly.TickerVol),
		TickerAnaly.TickerVol[0].CcyName,
	)
}
