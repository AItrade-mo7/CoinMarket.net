package global

import (
	"time"

	"CoinMarket.net/tmpl"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	// 初始化日志系统 保证日志可用
	mCycle.New(mCycle.Opt{
		Func:      LogInt,
		SleepTime: time.Hour * 8,
	}).Start()

	// 加载系统配置文件
	ServerEnvInt()

	// 加载用户配置文件
	UserInitEnv()

	logStr := `系统初始化完成`
	Log.Println(logStr)

	Email(EmailOpt{
		To:       UserEnv.Email.To,
		Subject:  "Start",
		Template: tmpl.Email,
		SendData: struct {
			Message string
			SysTime time.Time
			RunMod  int
		}{
			Message: "系统启动",
			SysTime: time.Now(),
			RunMod:  ServerEnv.RunMod,
		},
	}).Send()
}
