package dbTask

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EarnCountParam struct {
	StartTimeUnix int64  `bson:"StartTimeUnix"`
	EndTimeUnix   int64  `bson:"EndTimeUnix"`
	InstID        string `bson:"InstID"`
}

type EarnCountObj struct {
	CoinDB        *mMongo.DB // 币种数据
	AnalyDB       *mMongo.DB // 计算结果数据
	StartTimeUnix int64
	EndTimeUnix   int64
	KdataList     []mOKX.TypeKd
	AnalyList     []dbType.AnalyTickerType
}

func NewEarnCount(opt EarnCountParam) *EarnCountObj {
	var NewEarnObj EarnCountObj
	if opt.InstID == "BTC" || opt.InstID == "ETH" {
	} else {
		opt.InstID = "BTC"
	}

	Timeout := 4000 * 60
	NewEarnObj.AnalyDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("TickerAnaly")

	NewEarnObj.CoinDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection(opt.InstID + "USDT")

	NewEarnObj.StartTimeUnix = opt.StartTimeUnix
	NewEarnObj.EndTimeUnix = opt.StartTimeUnix

	return &NewEarnObj
}

func (_this *EarnCountObj) CountEnd() {
	_this.CoinDB.Close()
	_this.AnalyDB.Close()
}

func (_this *EarnCountObj) FindCoinKdata() *EarnCountObj {
	db := _this.CoinDB
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	findFK := bson.D{}
	findFK = append(findFK, bson.E{
		Key: "CreateTime",
		Value: bson.D{
			{
				Key:   "$gte", // 大于或等于
				Value: _this.StartTimeUnix,
			}, {
				Key:   "$lte", // 小于或等于
				Value: _this.EndTimeUnix,
			},
		},
	})

	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)

	var KdataList []mOKX.TypeKd

	for cursor.Next(db.Ctx) {
		var Kdata mOKX.TypeKd
		cursor.Decode(&Kdata)
		KdataList = append(KdataList, Kdata)

		global.Run.Println(Kdata.TimeStr, Kdata.InstID, Kdata.C)
	}

	if err != nil {
		global.LogErr(err)
		return nil
	}

	_this.KdataList = KdataList

	return _this
}

func (_this *EarnCountObj) FindAnaly() *EarnCountObj {
	db := _this.AnalyDB
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	findFK := bson.D{}
	findFK = append(findFK, bson.E{
		Key: "CreateTime",
		Value: bson.D{
			{
				Key:   "$gte", // 大于或等于
				Value: _this.StartTimeUnix,
			}, {
				Key:   "$lte", // 小于或等于
				Value: _this.EndTimeUnix,
			},
		},
	})

	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)

	var AnalyList []dbType.AnalyTickerType

	for cursor.Next(db.Ctx) {
		var AnalyTicker dbType.AnalyTickerType
		cursor.Decode(&AnalyTicker)
		AnalyList = append(AnalyList, AnalyTicker)

		global.Run.Println(AnalyTicker.TimeID, AnalyTicker.WholeDir)
	}

	if err != nil {
		global.LogErr(err)
		return nil
	}

	_this.AnalyList = AnalyList

	return _this
}
