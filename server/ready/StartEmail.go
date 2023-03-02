package ready

import (
	"os/exec"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
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

	taskPush.SysEmail(taskPush.SysEmailOpt{
		Subject:     "执行清理",
		Title:       "执行了一次清理",
		Message:     "执行一次" + config.File.ReClearShell + "脚本",
		Content:     mStr.ToStr(Succeed),
		Description: "系统清理",
	})
}

func SysReStart() {
	taskPush.SysEmail(taskPush.SysEmailOpt{
		Subject:     "Linux系统即将重启",
		Title:       "系统即将重启",
		Message:     "系统重启通知",
		Content:     "Linux系统 10 秒钟后 将自动进行重启",
		Description: "系统重启",
	})

	time.Sleep(time.Second * 10)

	isShellPath := mPath.Exists(config.File.SysReStart)
	if !isShellPath {
		global.Log.Println("未找到 SysReStart 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.SysReStart).Output()
	global.Log.Println("执行脚本", Succeed, err)
}
