package main

import (
	_ "embed"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/tmpl"
	"CoinMarket.net/server/utils/dbTidy"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)

	// 初始化系统参数
	global.Start()

	// ready.Start()

	// router.Start()

	inst.Start()
	dbTidy.FormatMarket()

	dbTidy.GetCoinKdata()

	go global.Email(global.EmailOpt{
		To: []string{
			"meichangliang@mo7.cc",
		},
		Subject:  "ServeStart",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "程序执行完毕",
			SysTime: time.Now(),
		},
	}).Send()
}
