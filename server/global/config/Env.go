package config

import (
	"fmt"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/spf13/viper"
)

var SysName = "CoinMarket.net"

var AppInfo struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
	Port    int    `bson:"Port"`
}

var SysEnv struct {
	MongoAddress   string
	MongoPassword  string
	MongoUserName  string
	MessageBaseUrl string
}

var BinanceKey = struct {
	ApiKey    string `bson:"ApiKey"`
	SecretKey string `bson:"SecretKey"`
}{
	ApiKey:    "e1ch7VN2DHqymTJMDRZLDjoyWHdjAaT2anp06elWJwOOZ51GoXXXQpOAeDenWEml",
	SecretKey: "E9w2QmhtJC6Z326mB9fcCFH6syB2qt7TbvwaFfmxWtrWdAFrQNTIVvT8bSKLtV9e",
}

func DefaultSysEnv() {
	SysEnv.MongoAddress = "tcy.mo7.cc:17017"
	SysEnv.MongoPassword = "mo7"
	SysEnv.MongoUserName = "asdasd55555"
	SysEnv.MessageBaseUrl = "http://msg.mo7.cc"
}

func LoadSysEnv(envPath string) {
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		DefaultSysEnv()
		errStr := fmt.Errorf("sys_env.yaml 读取配置文件出错，填充默认值 : %+v", err)
		LogErr(errStr)
	}
	viper.Unmarshal(&SysEnv)
}

func ServerEnvInit() {
	isLocalSysEnvFile := mPath.Exists(File.LocalSysEnv)
	isSysEnvFile := mPath.Exists(File.SysEnv)

	if isLocalSysEnvFile || isSysEnvFile {
		//
	} else {
		DefaultSysEnv()
		errStr := fmt.Errorf("没找到 sys_env.yaml 读取配置文件出错，填充默认值")
		LogErr(errStr)
	}

	if isLocalSysEnvFile {
		LoadSysEnv(File.LocalSysEnv)
	} else {
		LoadSysEnv(File.SysEnv)
	}
}
