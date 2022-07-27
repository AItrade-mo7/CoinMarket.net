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
	// 这里是启动日志

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

	// 获取 OKX 交易产品信息
	mCycle.New(mCycle.Opt{
		Func:      inst.Start,
		SleepTime: time.Hour * 8, // 每 8 时获取一次
	}).Start()
	go mFile.Write(config.Dir.JsonData+"/SPOT_inst.json", mJson.ToStr(okxInfo.SPOT_inst))
	go mFile.Write(config.Dir.JsonData+"/SWAP_inst.json", mJson.ToStr(okxInfo.SWAP_inst))

	mCycle.New(mCycle.Opt{
		Func:      GetTicker,
		SleepTime: time.Minute, // 每 1 分钟 获取一次
	}).Start()
	go mFile.Write(config.Dir.JsonData+"/TickerList.json", mJson.ToStr(okxInfo.TickerList))
}

func GetTicker() {
	binanceApi.GetTicker()
	tickers.GetTicker()
	SetTicker() // 在这里计算综合排行榜单

	TickerKdata()
}

/*


curl --request GET \
     --url https://api.exchange.coinbase.com/products/BTC-USDT/stats \
     --header 'Accept: application/json'

*/
