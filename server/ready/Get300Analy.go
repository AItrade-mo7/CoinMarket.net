package ready

import (
	"CoinMarket.net/server/global/apiType"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mMongo"
)

func GetTickerAnaly(opt dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
	resData = dbSearch.PagingType{}
	resErr = nil

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("TickerAnaly")
	defer db.Close()
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

		if opt.Type == "Client" {
			var ClientAnaly apiType.ClientAnalyType
			ClientAnaly.Unit = Analy.Unit
			ClientAnaly.TimeUnix = Analy.TimeUnix
			ClientAnaly.TimeStr = Analy.TimeStr
			ClientAnaly.TimeID = Analy.TimeID

			ClientAnaly.MaxUP = Analy.AnalyWhole[0].MaxUP.CcyName
			ClientAnaly.MaxUP_RosePer = Analy.AnalyWhole[0].MaxUP.RosePer

			ClientAnaly.MaxDown = Analy.AnalyWhole[0].MaxDown.CcyName
			ClientAnaly.MaxDown_RosePer = Analy.AnalyWhole[0].MaxDown.RosePer

		} else {
			AnyList = append(AnyList, Analy)
		}

	}

	// 生成返回数据
	resData = resCur.GenerateData(AnyList)

	return
}
