package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppInfo struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
	Port    int    `bson:"Port"`
}

var SysEnv struct {
	MongoAddress  string
	MongoPassword string
	MongoUserName string
	RunMod        int // 0 则为正常模式 ， 1 则为数据模拟模式
}

func LoadSysEnv(envPath string) {
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("sys_env.yaml 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&SysEnv)
}

type EmailInfo struct {
	Account  string
	Password string
	To       []string
}

var Email = EmailInfo{
	Account:  "trade@mo7.cc",
	Password: "Mcl931750",
	To: []string{
		"trade@mo7.cc",
	},
}
