package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var SysEnv struct {
	MongoAddress  string
	MongoPassword string
	MongoUserName string
}

func LoadSysEnv() {
	SysEnv.MongoAddress = "trade_api.mo7.cc:17017"
	SysEnv.MongoPassword = "asdasd55555"
	SysEnv.MongoUserName = "mo7"
}

var AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Port    int    `json:"Port"`
}

var AppEnv struct {
	RunMod int
}

func LoadAppEnv() {
	viper.SetConfigFile(File.AppEnv)

	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("AppEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&AppEnv)
}
