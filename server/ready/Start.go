package ready

import (
	"fmt"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnaly"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mMongo"
)

func Start() {
	// 发送启动邮件
	if config.AppEnv.RunMod == 0 {
		go global.Email(global.EmailOpt{
			To: []string{
				"meichangliang@mo7.cc",
			},
			Subject:  "ServeStart",
			Template: tmpl.SysEmail,
			SendData: tmpl.SysParam{
				Message: "服务启动",
				SysTime: time.Now(),
			},
		}).Send()
	}
	// 获取 OKX 交易产品信息
	mCycle.New(mCycle.Opt{
		Func:      inst.Start,
		SleepTime: time.Hour * 4, // 每 4 时获取一次
	}).Start()

	global.KdataLog.Println("okxInfo.SPOT_inst SWAP_inst", len(okxInfo.SPOT_inst), len(okxInfo.SWAP_inst))

	// 获取排行榜单
	mCycle.New(mCycle.Opt{
		Func:      GetTicker,
		SleepTime: time.Minute * 3, // 每 3 分钟获取一次
	}).Start()

	// 获取历史数据,并执行分析
	SetKdata("Start")
	go mClock.New(mClock.OptType{
		Func: TimerClickStart,
		Spec: "1 0,15,30,45 * * * ? ",
	})
}

func GetTicker() {
	binanceApi.GetTicker() //  获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合排行榜单    mOKX.TickerList  数据
}

// 获取历史数据

func TimerClickStart() {
	SetKdata("mClock")
}

func SetKdata(lType string) {
	TickerKdata()       // 获取并设置榜单币种最近 75 小时的历史数据 mOKX.MarketKdata   数据
	tickerAnaly.Start() // 开始对数据进行分析
	global.KdataLog.Println("okxInfo.TickerList ", len(okxInfo.TickerList), len(okxInfo.MarketKdata))

	if lType == "mClock" {
		SetMarketTickerDB()
	}
}

func SetMarketTickerDB() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()

	TickerDB := dbType.GetTickerDB()
	_, err := db.Table.InsertOne(db.Ctx, TickerDB)
	if err != nil {
		resErr := fmt.Errorf("注册,插入数据失败 %+v", err)
		global.LogErr(resErr)
		return
	}
}
