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

	// 数据榜单并进行数据库存储
	go mClock.New(mClock.OptType{
		Func: SetDBTickerData,
		Spec: "0,5,10,15,20,25,30,35,40,45,50,55 * * * * ? ",
	})
}

// 获取历史数据并存储
func SetDBTickerData() {
	time.Sleep(time.Second / 3)
	SetTickerAnaly() //  产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 以及 okxInfo.TickerAnaly
	go SetTickerAnalyDB()
	go SetCoinTickerDB()
	go SetCoinKdataDB("BTC")
	go SetCoinKdataDB("ETH")
}

// 获取榜单数据
func SetTickerAnaly() {
	binanceApi.GetTicker() // 获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据
	SetTickerKdata()       // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据
	global.Run.Println(
		"========= 开始分析 ===========",
		len(okxInfo.TickerVol),
		len(okxInfo.TickerKdata),
		len(okxInfo.TickerKdata["BTC-USDT"]),
	)
	okxInfo.TickerAnaly = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
		TickerVol:   okxInfo.TickerVol,
		TickerKdata: okxInfo.TickerKdata,
	})
}
