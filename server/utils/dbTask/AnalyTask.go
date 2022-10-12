package dbTask

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnalyTaskObj struct {
	TickerDB *mMongo.DB
	AnalyDB  *mMongo.DB
	CoinDB   *mMongo.DB
}

func NewAnalyTask() *AnalyTaskObj {
	var NewTask AnalyTaskObj

	Timeout := 4000 * 60
	dbTable := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect()
	NewTask.TickerDB = dbTable.Collection("CoinTicker")
	NewTask.AnalyDB = dbTable.Collection("TickerAnaly")
	NewTask.CoinDB = dbTable.Collection("BTCUSDT")

	return &NewTask
}

func (_this *AnalyTaskObj) CoinDBTraverse() {
	db := _this.CoinDB
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetSort(map[string]int{
		"TimeUnix": -1,
	})

	findFK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)
	for cursor.Next(db.Ctx) {
		var Kdata mOKX.TypeKd
		cursor.Decode(&Kdata)
		TimeID := mOKX.GetTimeID(Kdata.TimeUnix)

		global.Run.Println(TimeID, Kdata.InstID, Kdata.C)
	}

	if err != nil {
		global.LogErr(err)
	}
}

func StartEmail() {
	global.Email(global.EmailOpt{
		To: []string{
			"meichangliang@mo7.cc",
		},
		Subject:  "脚本执行",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "开始执行 AnalyTask",
			SysTime: time.Now(),
		},
	}).Send()
}

func EndEmail() {
	global.Email(global.EmailOpt{
		To: []string{
			"meichangliang@mo7.cc",
		},
		Subject:  "脚本结束",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "AnalyTask 执行完毕",
			SysTime: time.Now(),
		},
	}).Send()
}
