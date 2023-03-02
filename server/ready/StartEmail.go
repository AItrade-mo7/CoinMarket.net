package ready

import (
	"os/exec"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
)

func StartEmail() {
	err := taskPush.SysEmail(taskPush.SysEmailOpt{
		From:        config.SysName,
		Subject:     "系统启动",
		Title:       config.SysName + " 系统启动",
		Message:     "系统启动",
		Content:     mJson.Format(config.AppInfo),
		Description: "系统启动邮件",
	})
	global.Run.Println("系统启动邮件已发送", err)
}

func ReClearShell() {
	isShellPath := mPath.Exists(config.File.ReClearShell)
	if !isShellPath {
		global.Log.Println("未找到 ReClearShell 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.ReClearShell).Output()
	global.Log.Println("执行脚本", Succeed, err)

	// go global.Email(global.EmailOpt{
	// 	To:       config.Email.To,
	// 	Subject:  "数据库重启并执行清理",
	// 	Template: tmpl.SysEmail,
	// 	SendData: tmpl.SysParam{
	// 		Message: mStr.ToStr(Succeed),
	// 		SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
	// 	},
	// }).Send()
}

func SysReStart() {
	// go global.Email(global.EmailOpt{
	// 	To:       config.Email.To,
	// 	Subject:  "Linux系统即将重启",
	// 	Template: tmpl.SysEmail,
	// 	SendData: tmpl.SysParam{
	// 		Message: "Linux系统 10 秒钟后 将自动进行重启",
	// 		SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
	// 	},
	// }).Send()

	time.Sleep(time.Second * 10)

	isShellPath := mPath.Exists(config.File.SysReStart)
	if !isShellPath {
		global.Log.Println("未找到 SysReStart 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.SysReStart).Output()
	global.Log.Println("执行脚本", Succeed, err)
}
