package main

import (
	_ "embed"
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi"
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

	// MainTest()

	ready.Start()

	router.Start()

	// ==== 开始整理算法结果 ====
	// Task := dbTask.NewAnalyTask()
	// Task.CoinDBTraverse()

	// ==== 开始填充榜单历史 ====
	// FormatDB := dbTask.NewFormat()
	// FormatDB.TickerDBTraverse()

	// ==== 整理Kdata ====
	// dbTask.FormatKdata()
}

func MainTest() {
	okxApi.SetInst() // 获取并设置交易产品信息

	List := okxApi.GetKdata(okxApi.GetKdataOpt{
		InstID: "ETH-USDT",
		Size:   config.KdataLen,
		After:  1667740500000,
	})

	fmt.Println("List", len(List), List[len(List)-1].TimeStr)
}
