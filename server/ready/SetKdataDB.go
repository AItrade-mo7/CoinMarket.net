package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStruct"
	"go.mongodb.org/mongo-driver/bson"
)

func SetEthDB() {
	list := okxInfo.MarketKdata["ETH-USDT"]
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("ETHUSDT")
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
		var result mOKX.TypeKd
		db.Table.FindOne(db.Ctx, FK).Decode(&result)
		var err error
		lType := ""
		if result.TimeUnix > 0 {
			lType = "更新"
			_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
		} else {
			lType = "插入"
			_, err = db.Table.InsertOne(db.Ctx, Kd)
		}

		if err != nil {
			resErr := fmt.Errorf(lType+"数据失败 ETH %+v", err)
			global.LogErr(resErr)
		}
	}
}

func SetBtcDB() {
	list := okxInfo.MarketKdata["BTC-USDT"]

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("BTCUSDT")
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
		var result mOKX.TypeKd
		db.Table.FindOne(db.Ctx, FK).Decode(&result)
		var err error
		lType := ""
		if result.TimeUnix > 0 {
			lType = "更新"
			_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
		} else {
			lType = "插入"
			_, err = db.Table.InsertOne(db.Ctx, Kd)
		}

		if err != nil {
			resErr := fmt.Errorf(lType+"数据失败 BTC %+v", err)
			global.LogErr(resErr)
		}
	}
}

func SetMarketTickerDB() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()
}
