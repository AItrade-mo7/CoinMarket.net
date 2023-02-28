package global

import (
	"time"

	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mJson"
)

func Start() {
	// 初始化项目目录
	config.DirInit()

	// 初始化日志系统
	mCycle.New(mCycle.Opt{
		Func:      LogInit,
		SleepTime: time.Hour * 24,
	}).Start()

	// 加载 SysEnv
	config.ServerEnvInit()

	Log.Println(
		`系统初始化完成`,
		mJson.Format(config.Dir),
		mJson.Format(config.SysEnv),
		mJson.Format(config.AppInfo),
	)
}
