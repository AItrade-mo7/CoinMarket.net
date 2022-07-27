package ready

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
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
		SleepTime: time.Hour * 8, // 每 8 时获取一次
	}).Start()
	go mFile.Write(config.Dir.JsonData+"/SPOT_inst.json", mJson.ToStr(okxInfo.SPOT_inst))
	go mFile.Write(config.Dir.JsonData+"/SWAP_inst.json", mJson.ToStr(okxInfo.SWAP_inst))

	// 获取当前的行情与交易量榜单
	mCycle.New(mCycle.Opt{
		Func:      GetTicker,
		SleepTime: time.Minute * 5, // 每 5 分钟 获取一次
	}).Start()
	go mFile.Write(config.Dir.JsonData+"/TickerList.json", mJson.ToStr(okxInfo.TickerList))
	go mFile.Write(config.Dir.JsonData+"/MarketKdata.json", mJson.ToStr(okxInfo.MarketKdata))
}

func GetTicker() {
	binanceApi.GetTicker()
	tickers.GetTicker()
	SetTicker() // 在这里计算综合排行榜单
	TickerKdata()
}
