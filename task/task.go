package main

import (
	_ "embed"

	"CoinMarket.net/server/global"
)

func main() {
	// 初始化系统参数
	global.Start()
}
