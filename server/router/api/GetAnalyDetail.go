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
	TimeID string `bson:"TimeID"`
}

func GetAnalyDetail(c *fiber.Ctx) error {
	var json GetAnalyDetailParam
	mFiber.Parser(c, &json)

	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "CoinMarket",
	}).Connect()
	if err != nil {
		return c.JSON(result.ErrDB.WithData(err))
	}
	defer db.Close()
	db.Collection("TickerAnaly")

	FK := bson.D{{
		Key:   "TimeID",
		Value: json.TimeID,
	}}

	var Analy dbType.AnalyTickerType
	db.Table.FindOne(db.Ctx, FK).Decode(&Analy)

	return c.JSON(result.Succeed.WithData(Analy))
}
