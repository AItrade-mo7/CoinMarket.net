package dbTask

import (
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mMongo"
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
