package global

import (
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
)

func AppEnvInit() {
	// 检查配置文件在不在
	isUserEnvPath := mPath.Exists(config.File.AppEnv)
	if !isUserEnvPath {
		return
	}
	config.LoadAppEnv()
	Log.Println("加载 AppEnv : ", mJson.JsonFormat(mJson.ToJson(config.AppEnv)))
}
