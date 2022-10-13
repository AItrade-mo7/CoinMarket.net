package global

import (
	"time"

	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	// 初始化目录列表
	config.DirInit()

	// 初始化日志系统 保证日志可用
	mCycle.New(mCycle.Opt{
		Func:      LogInit,
		SleepTime: time.Hour * 8,
	}).Start()

	// 加载 SysEnv
	ServerEnvInit()

	// 发送启动邮件
	// if config.SysEnv.RunMod == 0 {
	// 	go Email(EmailOpt{
	// 		To: []string{
	// 			"meichangliang@mo7.cc",
	// 		},
	// 		Subject:  "ServeStart",
	// 		Template: tmpl.SysEmail,
	// 		SendData: tmpl.SysParam{
	// 			Message: "服务启动",
	// 			SysTime: mTime.IsoTime(false),
	// 		},
	// 	}).Send()
	// }
	Log.Println(`系统初始化完成`)
}
