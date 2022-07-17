package global

import (
	"flag"
	"fmt"

	"CoinMarket.net/config"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	viper "github.com/spf13/viper"
)

// 系统配置文件
var ServerEnv struct {
	LocalIP        string
	MongoAddress   string
	MongoPassword  string
	MongoUserName  string
	UpLoadFilePath string
	RunMod         int // 0 则为正常模式 ， 1 则为数据模拟模式
}

func LoadServerEnv(envPath string) {
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		LogErr(" ServerEnv 读取配置文件出错 ", err)
		return
	}
	viper.Unmarshal(&ServerEnv)
}

func ServerEnvInt() {
	Log.Println("\n",
		"AppPath", config.AppPath, "\n",
		"HomePath", config.HomePath, "\n",
		"LogPath", config.LogPath, "\n",
		"UserEnvPath", config.UserEnvPath, "\n",
		"StaticPath", config.StaticPath, "\n",
		"ClientFile", config.ClientFile, "\n",
		"HomeServerEnv", config.HomeServerEnv, "\n",
		"AppServerEnv", config.AppServerEnv,
	)

	isHomeEnvFile := mPath.Exists(config.HomeServerEnv)
	isAppEnvFile := mPath.Exists(config.AppServerEnv)

	if isHomeEnvFile {
		LoadServerEnv(config.HomeServerEnv)
	}
	if isAppEnvFile {
		LoadServerEnv(config.AppServerEnv)
	}

	if !isHomeEnvFile && !isAppEnvFile {
		errStr := fmt.Errorf(" 没找到 server_env.yaml 配置文件")
		LogErr(errStr)
		panic(errStr)
	}

	Log.Println("加载 ServerEnv : ", mStr.ToStr(ServerEnv))
}

// 用户配置文件
type EmailInfo struct {
	Account  string
	Password string
	To       []string
}

var UserEnv struct {
	Port  string
	Email EmailInfo
}

func LoadUserEnv() {
	viper.SetConfigFile(config.UserEnvPath)

	err := viper.ReadInConfig()
	if err != nil {
		LogErr(" UserEnv 读取配置文件出错 ", err)
		return
	}
	viper.Unmarshal(&UserEnv)
}

// 加载本地配置(设置默认值)
func UserInitEnv() {
	// 检查配置文件在不在
	isUserEnvPath := mPath.Exists(config.UserEnvPath)
	if isUserEnvPath {
		LoadUserEnv()
	}

	s := flag.String("port", UserEnv.Port, "端口号")
	flag.Parse()

	if len(*s) > 3 {
		// 存在参数则直接 设置为端口参数
		UserEnv.Port = *s
	}

	// 如果端口参数不存在,则设置为默认值参数
	if len(UserEnv.Port) < 3 {
		UserEnv.Port = "8998"
	}

	// 设置发件的 Email 信息
	UserEnv.Email.Account = "hunter_data_center@mo7.cc"
	UserEnv.Email.Password = "hIXY2pYSuxEz6Y5k"
	UserEnv.Email.To = []string{
		"meichangliang@mo7.cc",
	}

	Log.Println("加载 UserEnv : ", mStr.ToStr(UserEnv))
}
