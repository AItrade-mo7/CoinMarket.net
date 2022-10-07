package api

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type GetCoinListParam struct {
	InstID string `bson:"InstID"` // 只能为 BTC 或者 ETH
	dbSearch.FindParam
}

// 获取当前页码的币种数据，并进行存储，15分钟为限额
func GetCoinHistory(c *fiber.Ctx) error {
	var FetchParam GetCoinListParam
	mFiber.Parser(c, &FetchParam)

	var json dbSearch.FindParam
	jsoniter.Unmarshal(mJson.ToJson(FetchParam), &json)

	if FetchParam.InstID == "BTC" || FetchParam.InstID == "ETH" {
	} else {
		return c.JSON(result.Fail.WithData("InstID 只能为 BTC 或者 ETH"))
	}

	var err error
	var resData dbSearch.PagingType

	if FetchParam.InstID == "BTC" {
		resData, err = GetBTCKdata(json)
	}
	if FetchParam.InstID == "ETH" {
		resData, err = GetETHKdata(json)
	}
	if err != nil {
		return c.JSON(result.ErrDB.WithMsg(err))
	}

	return c.JSON(result.Succeed.WithData(resData))
}

func GetETHKdata(json dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("ETHUSDT")
	defer db.Close()
	err := db.Ping()
	if err != nil {
		db.Close()
		resErr = fmt.Errorf("数据读取失败,数据库连接错误1 %+v", err)
		global.LogErr(resErr)
		return
	}
	// 构建搜索参数
	resCur, err := dbSearch.GetCursor(dbSearch.CurOpt{
		Param: json,
		DB:    db,
	})
	if err != nil {
		resErr = err
		global.LogErr(resErr)
		return
	}

	// 提取数据
	var Kdata []any
	for resCur.Cursor.Next(db.Ctx) {
		var result mOKX.TypeKd
		resCur.Cursor.Decode(&result)

		Kdata = append(Kdata, result)
	}

	// 生成返回数据
	resData = resCur.GenerateData(Kdata)

	return
}

func GetBTCKdata(json dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("BTCUSDT")
	defer db.Close()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		db.Close()
		resErr = fmt.Errorf("数据读取失败,数据库连接错误1 %+v", err)
		global.LogErr(resErr)
		return
	}
	// 构建搜索参数
	resCur, err := dbSearch.GetCursor(dbSearch.CurOpt{
		Param: json,
		DB:    db,
	})
	if err != nil {
		resErr = err
		global.LogErr(resErr)
		return
	}

	// 提取数据
	var Kdata []any
	for resCur.Cursor.Next(db.Ctx) {
		var result mOKX.TypeKd
		resCur.Cursor.Decode(&result)
		Kdata = append(Kdata, result)
	}

	// 生成返回数据
	resData = resCur.GenerateData(Kdata)

	return
}
