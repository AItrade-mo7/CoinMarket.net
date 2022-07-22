package config

import (
	"os"

	"github.com/EasyGolang/goTools/mStr"
)

type DirType struct {
	App string // APP 根目录
	Log string // 日志文件目录
}

var Dir DirType

func DirInit() {
	Dir.App, _ = os.Getwd()

	Dir.Log = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"logs",
	)
}
