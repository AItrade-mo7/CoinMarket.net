package main

import (
	_ "embed"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi/kdata"
	"CoinMarket.net/server/ready"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)

	// 初始化系统参数
	global.Start()

	ready.Start()

	kdata.GetHistoryKdata(kdata.HistoryKdataParam{
		InstID:  "BTC-USDT",
		Current: 0,
	})

	// router.Start()
}
