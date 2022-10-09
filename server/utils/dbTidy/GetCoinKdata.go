package dbTidy

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStruct"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCoinKdata() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("CoinTicker")
	defer db.Close()

	FK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, FK)
	for cursor.Next(db.Ctx) {
		// 取出数据
		var curData map[string]any
		cursor.Decode(&curData)
		// 解析并格式化数据
		var Ticker dbType.CoinTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)

		// 当Kdata 数据不足时 请求 Kdata
		if len(Ticker.Kdata) < len(Ticker.TickerVol) {
			Ticker.Kdata = FetchKdata(Ticker.TickerVol)
		} else {
			// 数据充足则跳过
			return
		}

		// 查询Unix
		FK := bson.D{{
			Key:   "TimeUnix",
			Value: Ticker.TimeUnix,
		}}
		UK := bson.D{}
		mStruct.Traverse(Ticker, func(key string, val any) {
			UK = append(UK, bson.E{
				Key: "$set",
				Value: bson.D{
					{
						Key:   key,
						Value: val,
					},
				},
			})
		})
		// 解析结果
		var result dbType.CoinTickerTable
		db.Table.FindOne(db.Ctx, FK).Decode(&result)

		var err error
		lType := ""
		if result.TimeUnix > 0 {
			lType = "更新"
			_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
		} else {
			lType = "插入"
			_, err = db.Table.InsertOne(db.Ctx, Ticker)
		}

		if err != nil {
			resErr := fmt.Errorf(lType+"数据失败 ETH %+v", err)
			global.LogErr(resErr)
		}

	}

	if err != nil {
		resErr := fmt.Errorf("GetCoinKdata 失败 %+v", err)
		global.LogErr(resErr)
		return
	}
}

func FetchKdata(List []mOKX.TypeTicker) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)
	for _, val := range List {
		KdataList[val.InstID] = kdata.GetHistoryKdata(kdata.HistoryKdataParam{
			InstID: val.InstID,
			After:  val.Ts,
		})
	}
	return KdataList
}
