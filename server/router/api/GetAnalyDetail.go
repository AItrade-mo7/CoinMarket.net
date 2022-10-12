package api

import (
	"CoinMarket.net/server/global/apiType"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

type GetAnalyDetailParam struct {
	TimeID string `bson:"TimeID"`
}

func GetAnalyDetail(c *fiber.Ctx) error {
	var json GetAnalyDetailParam
	mFiber.Parser(c, &json)

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("CoinTicker")
	defer db.Close()

	FK := bson.D{
		{
			Key:   "TimeID",
			Value: json.TimeID,
		},
	}

	var curData map[string]any
	db.Table.FindOne(db.Ctx, FK).Decode(&curData)

	var returnData apiType.MarketTickerTable
	jsoniter.Unmarshal(mJson.ToJson(curData), &returnData)

	return c.JSON(result.Succeed.WithData(returnData))
}
