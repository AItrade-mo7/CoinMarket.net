package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	jsoniter "github.com/json-iterator/go"
)

func GetAnalyFirst300(json dbSearch.FindParam) (resData dbSearch.PagingType, resErr error) {
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
	var MarketTickerList []any
	for resCur.Cursor.Next(db.Ctx) {
		var result dbType.MarketTickerTable
		resCur.Cursor.Decode(&result)

		if json.Type == "Serve" {
			MarketTickerList = append(MarketTickerList, result)
		} else {
			var MarketTicker MarketTickerAPI
			jsoniter.Unmarshal(mJson.ToJson(result), &MarketTicker)
			MarketTicker.MaxUP = result.AnalyWhole[0].MaxUP.CcyName
			MarketTicker.MaxUP_RosePer = result.AnalyWhole[0].MaxUP.RosePer
			MarketTicker.MaxDown = result.AnalyWhole[0].MaxDown.CcyName
			MarketTicker.MaxDown_RosePer = result.AnalyWhole[0].MaxDown.RosePer
			MarketTickerList = append(MarketTickerList, MarketTicker)
		}
	}

	resData = resCur.GenerateData(MarketTickerList)

	return
}

type MarketTickerAPI struct {
	Unit            string `bson:"Unit"`
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
