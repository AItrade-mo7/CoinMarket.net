package main

import (
	_ "embed"
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi"
	"CoinMarket.net/server/okxApi/binanceApi"
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

	BinancePosition := binanceApi.GetAccount()

	fmt.Println(BinancePosition)

	// router.Start()

	// ==== 测试 ====
	// mClock.New(mClock.OptType{
	// 	Func: MainTest,
	// 	Spec: "10 0,5,10,15,20,25,30,35,40,45,50,55 * * * ? ", // 5 分的整数过 10 秒
	// })

	// ==== 开始整理算法结果 ====
	// dbTask.StartEmail()
	// Task := dbTask.NewAnalyTask()
	// Task.CoinDBTraverse()
	// dbTask.EndEmail("整理算法结果")

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
	})

	global.Run.Println("List", len(List), List[len(List)-1].TimeStr)
}
