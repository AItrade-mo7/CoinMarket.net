package dbTidy

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FormatMarket() {
	inst.Start()

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  3241 * 10000 * 60,
	}).Connect().Collection("MarketTicker")
	defer db.Close()

	dbCoin := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  3241 * 10000 * 60,
	}).Connect().Collection("CoinTicker")
	defer dbCoin.Close()

	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	FK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, FK, findOpt)

	for cursor.Next(db.Ctx) {
		var curData map[string]any
		cursor.Decode(&curData)
		var Ticker dbType.MarketTickerTable
		jsoniter.Unmarshal(mJson.ToJson(curData), &Ticker)

		// 开始带入新数据
		InsertCoinTicker(dbCoin, Ticker)
	}

	if err != nil {
		global.LogErr(err)
	}
}

func InsertCoinTicker(db *mMongo.DB, oldTicker dbType.MarketTickerTable) {
	// 用 oldTicker 构建 CoinTicker
	var CoinTickerData dbType.CoinTickerTable
	CoinTickerData.TickerVol = oldTicker.List
	CoinTickerData.TimeUnix = oldTicker.List[0].Ts
	CoinTickerData.TimeStr = mTime.UnixFormat(mStr.ToStr(oldTicker.List[0].Ts))

	FK := bson.D{{
		Key:   "TimeUnix",
		Value: CoinTickerData.TimeUnix,
	}}

	var dbTicker dbType.CoinTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&dbTicker)

	// 用榜单去请求 Kdata
	Ticker_Kdata := FetchKdata(CoinTickerData, dbTicker)

	if len(Ticker_Kdata) == 0 {
		global.Run.Println("跳过", dbTicker.TimeStr, len(dbTicker.Kdata), len(dbTicker.TickerVol))
		return
	}

	if len(Ticker_Kdata) != len(CoinTickerData.TickerVol) {
		global.LogErr("数据有问题", CoinTickerData.TimeStr, len(CoinTickerData.Kdata), len(CoinTickerData.TickerVol))
		return
	} else {
		CoinTickerData.Kdata = Ticker_Kdata
	}

	var err error
	lType := ""
	if dbTicker.TimeUnix > 0 {
		// 表示已经存在，则更新即可
		lType = "更新"
		UK := bson.D{}
		mStruct.Traverse(CoinTickerData, func(key string, val any) {
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
		_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
	} else {
		// 表示还未存在，则插入
		lType = "插入"
		_, err = db.Table.InsertOne(db.Ctx, CoinTickerData)
	}
	if err != nil {
		resErr := fmt.Errorf(lType+"数据失败 %+v", err, CoinTickerData.TimeUnix)
		global.LogErr(resErr)
		return
	}

	var newTicker dbType.CoinTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&newTicker)

	// if len(newTicker.Kdata) != len(newTicker.TickerVol) || len(newTicker.Kdata["BTC-USDT"]) < 280 {
	// 	global.LogErr("==错误==", lType, newTicker.TimeStr, len(newTicker.Kdata), len(newTicker.Kdata["BTC-USDT"]))
	// } else {
	global.Run.Println("==结束==", newTicker.TimeStr, len(newTicker.Kdata), len(newTicker.Kdata["BTC-USDT"]))
	// }
}

func FetchKdata(newTicker dbType.CoinTickerTable, dbTicker dbType.CoinTickerTable) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)

	for _, val := range newTicker.TickerVol {
		if len(dbTicker.Kdata[val.InstID]) > 290 {
			continue
		}
		// time.Sleep(time.Second / 3)
		// kdata_list := kdata.GetHistory300List(kdata.History300Param{
		// 	InstID: val.InstID,
		// 	After:  val.Ts,
		// })

		KdataList[val.InstID] = []mOKX.TypeKd{}
		global.Run.Println("请求结束", val.InstID, len(KdataList[val.InstID]))
	}

	return KdataList
}
