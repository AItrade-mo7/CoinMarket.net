package api

import (
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type GetAnalyDetailParam struct {
	CreateTimeUnix int64 `bson:"CreateTimeUnix"`
}

func GetAnalyDetail(c *fiber.Ctx) error {
	var json GetAnalyDetailParam
	mFiber.Parser(c, &json)

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()

	FK := bson.D{
		{
			Key:   "CreateTimeUnix",
			Value: json.CreateTimeUnix,
		},
	}

	var returnData dbType.MarketTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&returnData)

	return c.JSON(result.Succeed.WithData(returnData))
}
