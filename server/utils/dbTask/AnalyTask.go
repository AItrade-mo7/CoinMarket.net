package dbTask

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi"
	"CoinMarket.net/server/tickerAnaly"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnalyTaskObj struct {
	TickerDB *mMongo.DB
	AnalyDB  *mMongo.DB
	CoinDB   *mMongo.DB
}

func NewAnalyTask() *AnalyTaskObj {
	okxApi.SetInst() // 获取并设置交易产品信息

	var NewTask AnalyTaskObj

	Timeout := 4000 * 60

	NewTask.TickerDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("CoinTicker")

	NewTask.AnalyDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("TickerAnaly")

	NewTask.CoinDB = mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection("BTCUSDT")

	return &NewTask
}

func (_this *AnalyTaskObj) CoinDBTraverse() {
	db := _this.CoinDB
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})

	findFK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	for cursor.Next(db.Ctx) {
		var Kdata mOKX.TypeKd
		cursor.Decode(&Kdata)
		_this.FindTicker(Kdata)
	}

	if err != nil {
		global.LogErr(err)
	}
}

func (_this *AnalyTaskObj) FindTicker(Kdata mOKX.TypeKd) {
	db := _this.TickerDB

	FK := bson.D{{
		Key:   "TimeID",
		Value: mOKX.GetTimeID(Kdata.TimeUnix),
	}}

	var Ticker dbType.CoinTickerTable
	db.Table.FindOne(db.Ctx, FK).Decode(&Ticker)

	BtcList := Ticker.Kdata["BTC-USDT"]
	if len(BtcList) > 90 && len(Ticker.TickerVol) == len(Ticker.Kdata) && len(Ticker.TickerVol) > 3 {
	} else {
		Ticker.TimeID = mOKX.GetTimeID(Kdata.TimeUnix)
		Ticker.TimeUnix = Kdata.TimeUnix
		Ticker.TimeStr = Kdata.TimeStr
		global.Log.Println("数据为空,造一个数据", Ticker.TimeID)
	}
	_this.AnalyStart(Ticker)
}

func (_this *AnalyTaskObj) AnalyStart(Ticker dbType.CoinTickerTable) {
	db := _this.AnalyDB

	BtcList := Ticker.Kdata["BTC-USDT"]

	AnalyResult := dbType.AnalyTickerType{}
	if len(BtcList) > 90 && len(Ticker.TickerVol) > 3 && len(Ticker.TickerVol) == len(Ticker.Kdata) {
		AnalyResult = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
			TickerVol:   Ticker.TickerVol,
			TickerKdata: Ticker.Kdata,
		})
	} else {
		global.Run.Println(
			"数据错误",
			Ticker.TimeID,
			len(Ticker.TickerVol),
			len(Ticker.Kdata),
			len(BtcList),
		)
		AnalyResult.Unit = config.Unit
		AnalyResult.TimeUnix = Ticker.TimeUnix
		AnalyResult.TimeStr = Ticker.TimeStr
		AnalyResult.TimeID = Ticker.TimeID
	}

	FK := bson.D{{
		Key:   "TimeID",
		Value: AnalyResult.TimeID,
	}}

	UK := bson.D{}
	mStruct.Traverse(AnalyResult, func(key string, val any) {
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
	upOpt := options.Update()
	upOpt.SetUpsert(true)
	_, err := db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
	if err != nil {
		global.LogErr("数据更插失败", err)
	}

	global.Run.Println(
		"更插完毕",
		AnalyResult.TimeID,
		len(AnalyResult.TickerVol),
	)
}

func StartEmail() {
	go global.Email(global.EmailOpt{
		To:       config.Email.To,
		Subject:  "脚本执行",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "开始执行 AnalyTask",
			SysTime: mTime.IsoTime(),
		},
	}).Send()
	global.Run.Println("======= 脚本开始 =======")
}

func EndEmail(Msg string) {
	global.Run.Println("======= 脚本结束 =======")

	lMsg := Msg

	if len(lMsg) < 1 {
		lMsg = "执行完毕"
	}

	global.Email(global.EmailOpt{
		To:       config.Email.To,
		Subject:  "脚本结束",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: lMsg,
			SysTime: mTime.IsoTime(),
		},
	}).Send()
}
