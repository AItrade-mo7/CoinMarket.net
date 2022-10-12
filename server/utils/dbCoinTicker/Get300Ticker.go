package dbCoinTicker

import (
	"log"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/tickerAnaly"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mMongo"
)

func GetTickerList(json dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
	resData = dbSearch.PagingType{}
	resErr = nil
	log.Println("开始读取")

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  3241 * 10000 * 60,
	}).Connect().Collection("CoinTicker")
	defer db.Close()

	// 构建搜索参数
	resCur, err := dbSearch.GetCursor(dbSearch.CurOpt{
		Param: json,
		DB:    db,
	})
	if err != nil {
		resErr = err
		return
	}
	// 提取数据
	var ResultList []any
	for resCur.Cursor.Next(db.Ctx) {
		var Ticker dbType.CoinTickerTable
		resCur.Cursor.Decode(&Ticker)
		// 该过程2秒钟
		AnalyResult := tickerAnaly.GetAnaly(tickerAnaly.TickerAnalyParam{
			TickerVol:   Ticker.TickerVol,
			TickerKdata: Ticker.Kdata,
		})
		ResultList = append(ResultList, AnalyResult)
	}
	// 生成返回数据
	resData = resCur.GenerateData(ResultList)

	return
}
