package ready

import (
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
)

func Start() {
	// 发送启动邮件
	if config.SysEnv.RunMod == 0 {
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
	global.KdataLog.Println("ready.Start inst.Start", len(okxInfo.SPOT_inst), len(okxInfo.SWAP_inst))

	// 获取排行榜单
	mCycle.New(mCycle.Opt{
		Func:      SetTickerAnaly,
		SleepTime: time.Minute * 5, // 每 5 分钟获取一次
	}).Start()

	// 获取历史数据,并执行分析
	SetKdata("Start")
	go mClock.New(mClock.OptType{
		Func: TimerClickStart,
		Spec: "0 0,15,30,45 * * * ? ",
	})
}

func SetTickerAnaly() {
	binanceApi.GetTicker() //  获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合排行榜单 产出 okxInfo.TickerVol 数据
	SetTickerKdata()       // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据

	global.Run.Println(
		"========= 开始分析 ===========",
		len(okxInfo.TickerVol),
		okxInfo.TickerVol[0].CcyName,
		len(okxInfo.TickerKdata),
		len(okxInfo.TickerKdata["BTC-USDT"]),
	)

	// 在这里计算分析结果
	okxInfo.TickerAnaly = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
		TickerVol:   okxInfo.TickerVol,
		TickerKdata: okxInfo.TickerKdata,
	})
}

// 获取历史数据

func TimerClickStart() {
	time.Sleep(time.Second)
	SetKdata("mClock")
}

func SetKdata(lType string) {
	SetTickerAnaly() //  产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 以及 okxInfo.TickerAnaly

	if lType == "mClock" {
		go SetTickerAnalyDB()
		go SetCoinTickerDB()
		go SetCoinKdataDB("BTC")
		go SetCoinKdataDB("ETH")
	}
}
