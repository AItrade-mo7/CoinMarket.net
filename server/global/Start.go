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
		SleepTime: time.Hour * 24,
	}).Start()

	// 加载 SysEnv
	ServerEnvInit()

	// Log.Println(`系统初始化完成`, mJson.Format(config.SysEnv))
}
