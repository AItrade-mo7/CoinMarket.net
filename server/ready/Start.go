package ready

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/tickerAnalyse"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	if config.AppEnv.RunMod == 0 {
		// 发送邮件
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
		SleepTime: time.Hour * 16, // 每 16 时获取一次
	}).Start()

	// 获取当前的行情与交易量榜单
	mCycle.New(mCycle.Opt{
		Func:      GetTicker,
		SleepTime: time.Minute * 5, // 每 5 分钟 获取一次
	}).Start()

	// ana := tickerAnalyse.Single["ETC-USDT"]
	// fmt.Println("ana", ana.DiffHour)
}

func GetTicker() {
	binanceApi.GetTicker() //  获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合排行榜单    okxInfo.TickerList  数据
	TickerKdata()          // 获取并设置榜单币种最近 75 小时的历史数据 okxInfo.MarketKdata   数据

	tickerAnalyse.Start() // 开始对数据进行分析
}
