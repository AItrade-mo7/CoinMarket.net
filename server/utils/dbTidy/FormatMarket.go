package dbTidy

import (
	"fmt"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

func FormatMarket() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  3189 * 10 * 10,
	}).Connect().Collection("MarketTicker")
	defer db.Close()

	dbCoin := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  3189 * 10 * 10,
	}).Connect().Collection("CoinTicker")
	defer dbCoin.Close()

	FK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, FK)

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
	// 用榜单去请求 Kdata
	Ticker_Kdata := FetchKdata(CoinTickerData)

	FK := bson.D{{
		Key:   "TimeUnix",
		Value: CoinTickerData.TimeUnix,
	}}
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

	var Ticker dbType.CoinTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&Ticker)

	if len(Ticker_Kdata) == 0 {
		global.Run.Println("跳过", Ticker.TimeStr, len(Ticker.Kdata), len(Ticker.TickerVol))
		return
	}

	if len(Ticker_Kdata) != len(CoinTickerData.TickerVol) {
		global.Run.Println("数据有问题", Ticker.TimeStr, len(Ticker.Kdata), len(Ticker.TickerVol))
		return
	} else {
		CoinTickerData.Kdata = Ticker_Kdata
	}

	var err error
	lType := ""
	if Ticker.TimeUnix > 0 {
		// 表示已经存在，则更新即可
		lType = "更新"
		global.Run.Println("进行数据更新", CoinTickerData.TimeStr)
		_, err = db.Table.UpdateOne(db.Ctx, FK, UK)
	} else {
		// 表示还未存在，则插入
		lType = "插入"
		global.Run.Println("进行数据插入", CoinTickerData.TimeStr)
		_, err = db.Table.InsertOne(db.Ctx, CoinTickerData)
	}
	if err != nil {
		resErr := fmt.Errorf(lType+"数据失败 %+v", err, CoinTickerData.TimeUnix)
		global.LogErr(resErr)
		return
	}

	global.Run.Println("====结束======", Ticker.TimeStr, len(Ticker.Kdata), len(Ticker.Kdata["BTC-USDT"]))
}

func FetchKdata(Ticker dbType.CoinTickerTable) map[string][]mOKX.TypeKd {
	KdataList := make(map[string][]mOKX.TypeKd)

	for _, val := range Ticker.TickerVol {
		if len(Ticker.Kdata[val.InstID]) < 280 {
			kdata_list := []mOKX.TypeKd{}

			time.Sleep(time.Second / 8)
			kdata_1 := kdata.GetHistoryKdata(kdata.HistoryKdataParam{
				InstID:  val.InstID,
				Current: 0,
				After:   val.Ts,
				Size:    100,
			})
			kdata_list = append(kdata_list, kdata_1...)
			time.Sleep(time.Second / 8)
			kdata_2 := kdata.GetHistoryKdata(kdata.HistoryKdataParam{
				InstID:  val.InstID,
				Current: 1,
				After:   val.Ts,
				Size:    100,
			})
			kdata_list = append(kdata_list, kdata_2...)
			time.Sleep(time.Second / 8)
			kdata_3 := kdata.GetHistoryKdata(kdata.HistoryKdataParam{
				InstID:  val.InstID,
				Current: 2,
				After:   val.Ts,
				Size:    100,
			})
			kdata_list = append(kdata_list, kdata_3...)

			KdataList[val.InstID] = kdata_list

			global.Run.Println("请求结束", val.InstID, len(kdata_list))

		}
	}

	return KdataList
}
