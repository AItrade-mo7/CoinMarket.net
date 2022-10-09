package ready

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/okxInfo"
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
		Func:      GetTicker,
		SleepTime: time.Minute * 5, // 每 5 分钟获取一次
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
	global.Run.Println("========= 开始获取榜单数据 ===========")
	GetTicker()

	global.Run.Println("开始获取历史价格")
	TickerKdata() // 获取并设置榜单币种最近 75 小时的历史数据 mOKX.MarketKdata   数据

	if lType == "mClock" {
		SetMarketTickerDB()
		SetEthDB()
		SetBtcDB()
	}
}
