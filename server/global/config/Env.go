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
	SysEnv.MongoAddress = "aws.mo7.cc:17017"
	SysEnv.MongoPassword = "asdasd55555"
	SysEnv.MongoUserName = "mo7"
}

var AppInfo struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
	Port    int    `bson:"Port"`
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
