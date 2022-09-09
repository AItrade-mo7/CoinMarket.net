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

func GetAnalyHistory(c *fiber.Ctx) error {
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
		var result map[string]any
		resCur.Cursor.Decode(&result)

		var MarketTicker dbType.MarketTickerTable
		jsoniter.Unmarshal(mJson.ToJson(result), &MarketTicker)

		MarketTickerList = append(MarketTickerList, MarketTicker)
	}

	returnData := resCur.GenerateData(MarketTickerList)

	mJson.Println(json)

	return c.JSON(result.Succeed.WithData(returnData))
}
