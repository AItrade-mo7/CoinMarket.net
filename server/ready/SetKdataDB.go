package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetEthDB() {
	list := okxInfo.TickerKdata["ETH-USDT"]
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("ETHUSDT")
	defer global.Run.Println("关闭数据库连接 ETHUSDT")
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
			resErr := fmt.Errorf("数据更插失败 ETH %+v", err)
			global.LogErr(resErr)
		}
	}
}

func SetBtcDB() {
	list := okxInfo.TickerKdata["BTC-USDT"]

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("BTCUSDT")
	defer global.Run.Println("关闭数据库连接 BTCUSDT")
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
			resErr := fmt.Errorf("数据更插失败 BTC %+v", err)
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
	}).Connect().Collection("CoinTicker")
	defer global.Run.Println("关闭数据库连接 CoinTicker")
	defer db.Close()

	CoinTickerData := dbType.JoinNowCoinTicker(okxInfo.TickerList, okxInfo.TickerKdata)

	FK := bson.D{{
		Key:   "TimeID",
		Value: CoinTickerData.TimeID,
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

	upOpt := options.Update()
	upOpt.SetUpsert(true)
	_, err := db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
	if err != nil {
		resErr := fmt.Errorf("数据更插失败 Ticker %+v", err)
		global.LogErr(resErr)
	}

	var newTicker dbType.CoinTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&newTicker)

	global.Run.Println("CoinTicker", "更插完毕", newTicker.TimeStr, len(newTicker.TickerVol), len(newTicker.Kdata))
}

func DBClose(db *mMongo.DB) {
	global.Run.Println("数据库关闭链接")
	db.Close()
}
