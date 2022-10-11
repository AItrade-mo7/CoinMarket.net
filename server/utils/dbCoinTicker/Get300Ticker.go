package dbCoinTicker

import (
	"fmt"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mMongo"
)

func GetTickerList(json dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
	resData = dbSearch.PagingType{}
	resErr = nil

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
	var CoinTickerList []any
	for resCur.Cursor.Next(db.Ctx) {
		var Ticker dbType.CoinTickerTable
		resCur.Cursor.Decode(&Ticker)
		CoinTickerList = append(CoinTickerList, Ticker)

		fmt.Println(Ticker.TimeID, len(Ticker.TickerVol), len(Ticker.Kdata), len(Ticker.Kdata["BTC-USDT"]))
	}

	return
}
