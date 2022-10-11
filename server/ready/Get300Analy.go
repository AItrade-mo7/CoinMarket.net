package ready

import (
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mMongo"
)

func Get300Analy(opt dbSearch.FindParam) (resData []any, resErr error) {
	resData = []any{}
	resErr = nil

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("CoinTicker")
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
	// var AnyList []any
	for resCur.Cursor.Next(db.Ctx) {
		var result map[string]any
		resCur.Cursor.Decode(&result)

		// var AbilityActive dbType.AbilityActiveTable
		// jsoniter.Unmarshal(mJson.ToJson(result), &AbilityActive)

		// abilityActiveList = append(abilityActiveList, AbilityActive)
	}

	// 生成返回数据
	// returnData := resCur.GenerateData(AnyList)

	return
}
