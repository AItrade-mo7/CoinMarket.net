package main

import (
	_ "embed"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/ready"
	"CoinMarket.net/server/router"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)

	// 初始化系统参数
	global.Start()

	ready.Start()

	router.Start()

	// select {} // 阻塞主线程停止

	// dbTidy.RemovalDup()
}
