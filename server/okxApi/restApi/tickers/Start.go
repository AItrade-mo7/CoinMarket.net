package tickers

import (
	"CoinMarket.net/server/okxApi/binanceApi"
)

// 获取行情信息

func Start() {
	binanceApi.GetTicker()
	GetTicker()
}
