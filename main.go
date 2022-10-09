package main

import (
	_ "embed"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
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

	global.Email(global.EmailOpt{
		To: []string{
			"meichangliang@mo7.cc",
		},
		Subject:  "ServeStart",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "开始执行脚本",
			SysTime: time.Now(),
		},
	}).Send()

	dbTidy.FormatMarket()

	global.Email(global.EmailOpt{
		To: []string{
			"meichangliang@mo7.cc",
		},
		Subject:  "ServeStart",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "脚本执行结束",
			SysTime: time.Now(),
		},
	}).Send()
}
