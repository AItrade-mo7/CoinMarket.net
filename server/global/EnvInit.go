package global

import (
	"fmt"

	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mPath"
)

func ServerEnvInit() {
	isLocalSysEnvFile := mPath.Exists(config.File.LocalSysEnv)
	isSysEnvFile := mPath.Exists(config.File.SysEnv)

	if isLocalSysEnvFile || isSysEnvFile {
		//
	} else {
		errStr := fmt.Errorf("没找到 sys_env.yaml 配置文件")
		LogErr(errStr)
		panic(errStr)
	}

	if isLocalSysEnvFile {
		config.LoadSysEnv(config.File.LocalSysEnv)
	} else {
		config.LoadSysEnv(config.File.SysEnv)
	}
}
