package ready

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
)

func Start() {
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
}

/*


curl --request GET \
     --url https://api.exchange.coinbase.com/products/BTC-USDT/stats \
     --header 'Accept: application/json'

*/
