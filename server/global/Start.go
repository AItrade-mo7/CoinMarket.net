package global

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTime"
)

func Start() {
	// 初始化目录列表
	config.DirInit()

	// 初始化日志系统 保证日志可用
	mCycle.New(mCycle.Opt{
		Func:      LogInit,
		SleepTime: time.Hour * 24,
	}).Start()

	// 加载 SysEnv
	ServerEnvInit()

	// 发送启动邮件
	go Email(EmailOpt{
		To: []string{
			"meichangliang@mo7.cc",
		},
		Subject:  "ServeStart",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "服务启动",
			SysTime: mTime.IsoTime(),
		},
	}).Send()
	Log.Println(`系统初始化完成`, mJson.Format(config.SysEnv))
}
