package api

import (
	"CoinMarket.net/server/global/apiType"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/gofiber/fiber/v2"
)

func GetAnalyList(c *fiber.Ctx) error {
	var json dbSearch.FindParam
	mFiber.Parser(c, &json)

	resData, err := GetTickerAnaly(json)
	if err != nil {
		return c.JSON(result.Fail.WithData(err))
	}

	return c.JSON(result.Succeed.WithData(resData))
}

func GetTickerAnaly(opt dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
	resData = dbSearch.PagingType{}
	resErr = nil

	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "CoinMarket",
	}).Connect()
	if err != nil {
		resErr = err
		return
	}
	defer db.Close()
	db.Collection("TickerAnaly")
	// 构建搜索参数

	resCur, err := dbSearch.GetCursor(dbSearch.CurOpt{
		Param: opt,
		DB:    db,
	})
	if err != nil {
		resErr = err
		return
	}

	// 提取数据
	var AnyList []any
	for resCur.Cursor.Next(db.Ctx) {
		var Analy dbType.AnalyTickerType
		resCur.Cursor.Decode(&Analy)

		if opt.Type == "Serve" {
			AnyList = append(AnyList, Analy)
		} else {
			var ClientAnaly apiType.ClientAnalyType
			ClientAnaly.Unit = Analy.Unit
			ClientAnaly.TimeUnix = Analy.TimeUnix
			ClientAnaly.TimeStr = Analy.TimeStr
			ClientAnaly.TimeID = Analy.TimeID
			ClientAnaly.Version = Analy.Version
			ClientAnaly.MillionCoin = Analy.MillionCoin
			if len(Analy.AnalyWhole) > 0 {
				ClientAnaly.MaxUP = Analy.AnalyWhole[0].MaxUP.CcyName
				ClientAnaly.MaxUP_RosePer = Analy.AnalyWhole[0].MaxUP.RosePer
				ClientAnaly.MaxDown = Analy.AnalyWhole[0].MaxDown.CcyName
				ClientAnaly.MaxDown_RosePer = Analy.AnalyWhole[0].MaxDown.RosePer
			}

			AnyList = append(AnyList, ClientAnaly)

		}

	}

	// 生成返回数据
	resData = resCur.GenerateData(AnyList)

	return
}
