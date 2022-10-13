package main

import (
	_ "embed"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/ready"
	"CoinMarket.net/server/router"
	"CoinMarket.net/server/utils/dbTask"
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

	// ==== 开始整理算法结果 ====
	dbTask.StartEmail()
	Task := dbTask.NewAnalyTask()
	Task.CoinDBTraverse()
	dbTask.EndEmail()
}
