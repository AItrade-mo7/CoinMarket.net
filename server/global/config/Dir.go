package config

import (
	"os"

	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

type DirType struct {
	Home     string // Home 根目录
	App      string // APP 根目录
	Log      string // 日志文件目录
	JsonData string // json 数据存放目录
}

var Dir DirType

type FileType struct {
	SysEnv       string // /root/sys_env.yaml
	LocalSysEnv  string // ./sys_env.yaml
	ReClearShell string // 清理系统缓存的脚本
	SysReStart   string // 系统重启脚本
}

var File FileType

func DirInit() {
	Dir.Home = mPath.HomePath()

	Dir.App, _ = os.Getwd()

	Dir.Log = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"logs",
	)

	Dir.JsonData = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"jsonData",
	)

	File.SysEnv = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"sys_env.yaml",
	)
	File.LocalSysEnv = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"sys_env.yaml",
	)

	File.ReClearShell = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"ReClear.sh",
	)

	File.SysReStart = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"SysReStart.sh",
	)

	if !mPath.Exists(File.ReClearShell) {
		mFile.Write(File.ReClearShell, tmpl.ReClear)
	}

	if !mPath.Exists(File.SysReStart) {
		mFile.Write(File.SysReStart, tmpl.SysReStart)
	}
}
