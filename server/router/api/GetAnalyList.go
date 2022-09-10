package api

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func GetAnalyList(c *fiber.Ctx) error {
	var json dbSearch.FindParam
	mFiber.Parser(c, &json)

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()
	err := db.Ping()
	if err != nil {
		db.Close()
		errStr := fmt.Errorf("数据读取失败,数据库连接错误1 %+v", err)
		global.LogErr(errStr)
		return c.JSON(result.ErrDB.WithMsg(errStr))
	}

	// 构建搜索参数
	resCur, err := dbSearch.GetCursor(dbSearch.CurOpt{
		Param: json,
		DB:    db,
		Lang:  c.Get("Lang"),
	})
	if err != nil {
		return c.JSON(result.ErrDB.WithMsg(err))
	}

	// 提取数据
	var MarketTickerList []any
	for resCur.Cursor.Next(db.Ctx) {
		var result dbType.MarketTickerTable
		resCur.Cursor.Decode(&result)

		var MarketTicker MarketTickerAPI
		jsoniter.Unmarshal(mJson.ToJson(result), &MarketTicker)
		MarketTicker.MaxUP = result.AnalyWhole[0].MaxUP.CcyName
		MarketTicker.MaxUP_RosePer = result.AnalyWhole[0].MaxUP.RosePer
		MarketTicker.MaxDown = result.AnalyWhole[0].MaxDown.CcyName
		MarketTicker.MaxDown_RosePer = result.AnalyWhole[0].MaxDown.RosePer

		MarketTickerList = append(MarketTickerList, MarketTicker)
	}

	returnData := resCur.GenerateData(MarketTickerList)

	return c.JSON(result.Succeed.WithData(returnData))
}

type MarketTickerAPI struct {
	WholeDir        int    `bson:"WholeDir"`
	TimeUnix        int64  `bson:"TimeUnix"`
	Time            string `bson:"Time"`
	CreateTimeUnix  int64  `bson:"CreateTimeUnix"`
	CreateTime      string `bson:"CreateTime"`
	MaxUP           string `json:"MaxUP"` // 最大涨幅币种
	MaxUP_RosePer   string `json:"MaxUP_RosePer"`
	MaxDown         string `json:"MaxDown"` // 最大跌幅币种
	MaxDown_RosePer string `json:"MaxDown_RosePer"`
}
