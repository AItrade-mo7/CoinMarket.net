package tickers

import "CoinMarket.net/server/okxApi/okxInfo"

// 获取行情信息

var OKXTickerList []okxInfo.TickerType // okx的Ticker

func Start() {
	GetTicker()
}
