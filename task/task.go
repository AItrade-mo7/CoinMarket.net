package main

import (
	_ "embed"

	"CoinMarket.net/server/global"
	"CoinMarket.net/task/dbTask"
)

func main() {
	// 初始化系统参数
	global.Start()

	// 新建回测
	// back := testHunter.New(testHunter.TestOpt{
	// 	StartTime: mTime.TimeParse(mTime.Lay_ss, "2022-12-12T00:00:00"),
	// 	EndTime:   mTime.TimeParse(mTime.Lay_ss, "2023-01-01T00:00:00"),
	// 	InstID:    "BTC-USDT",
	// })
	// back.StuffDBKdata(func(KD mOKX.TypeKd) {
	// 	global.Run.Println(KD.TimeStr, KD.C)
	// })

	dbTask.FormatKdata()
}
