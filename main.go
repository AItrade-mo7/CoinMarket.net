package main

import (
	"CoinMarket.net/global"
	"CoinMarket.net/observe"
	"CoinMarket.net/router"
)

// https://juejin.cn/post/6987204577879654407

func main() {
	// // 初始化系统参数
	global.Start()

	// // 数据整备
	observe.Start()

	// 启动端口服务
	router.Start()
}
