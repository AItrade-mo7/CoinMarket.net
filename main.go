package main

import (
	_ "embed"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/utils/dbTask"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)
	// 初始化系统参数
	global.Start()

	// 子任务
	Task()

	// 数据准备
	// ready.Start()

	// 启动 http 监听服务
	// router.Start()
}

func Task() {
	// 回测系统
	dbTask.BackTest(dbTask.BackTestOpt{
		StartTime: "2023-1-1",
		EndTime:   "2023-2-1",
	})
}
