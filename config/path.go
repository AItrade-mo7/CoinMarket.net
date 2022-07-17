package config

import (
	"os"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

// 当前目录
var AppPath, _ = os.Getwd()

// 系统根目录
var HomePath = mPath.HomePath()

// 日志文件目录
var LogPath = mStr.Join(
	AppPath,
	mStr.ToStr(os.PathSeparator),
	"logs",
)

// 用户配置文件

var UserEnvPath = mStr.Join(
	AppPath,
	mStr.ToStr(os.PathSeparator),
	"user_config.yaml",
)

// 静态资源目录
var StaticPath = mStr.Join(
	AppPath,
	mStr.ToStr(os.PathSeparator),
	"www",
)

// 客户端文件
var ClientFile = mStr.Join(
	StaticPath,
	mStr.ToStr(os.PathSeparator),
	"index.html",
)

// 系统配置文件

var HomeServerEnv = mStr.Join(
	HomePath,
	mStr.ToStr(os.PathSeparator),
	"server_env.yaml",
)

var AppServerEnv = mStr.Join(
	AppPath,
	mStr.ToStr(os.PathSeparator),
	"server_env.yaml",
)

var SecretKey = mEncrypt.MD5("golang is good")
