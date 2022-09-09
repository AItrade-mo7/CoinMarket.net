package api

import (
	"fmt"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetAnalyHistoryParam struct {
	Current int64 `json:"Current"` // 列表
}

func GetAnalyHistory(c *fiber.Ctx) error {
	var json GetAnalyHistoryParam
	mFiber.Parser(c, &json)

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()

	FK := bson.D{}
	findOpt := options.Find()

	findOpt.SetLimit(300)
	findOpt.SetSort(bson.D{{
		Key:   "TimeUnix",
		Value: -1,
	}})
	findOpt.SetSkip(json.Current * 300)

	cur, err := db.Table.Find(db.Ctx, FK, findOpt)
	if err != nil {
		db.Close()
		resErr := fmt.Errorf("数据读取失败 %+v", err)
		return c.JSON(result.ErrDB.WithData(resErr))
	}

	var MarketTickerList []dbType.MarketTickerTable
	for cur.Next(db.Ctx) {
		var result dbType.MarketTickerTable
		cur.Decode(&result)
		MarketTickerList = append(MarketTickerList, result)
	}

	return c.JSON(result.Succeed.WithData(MarketTickerList))
}
