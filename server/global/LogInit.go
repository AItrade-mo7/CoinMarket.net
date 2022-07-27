package global

import (
	"fmt"
	"log"
	"os"

	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mLog"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTime"
)

var (
	Log      *log.Logger // 系统日志& 重大错误或者事件
	InstLog  *log.Logger // 产品信息接口报错
	KdataLog *log.Logger // Rest 接口的请求日志
)

func LogInit() {
	// 检测 logs 目录
	isLogPath := mPath.Exists(config.Dir.Log)
	if !isLogPath {
		// 不存在则创建 logs 目录
		os.MkdirAll(config.Dir.Log, 0o777)
	}

	isJsonDataPath := mPath.Exists(config.Dir.JsonData)
	if !isJsonDataPath {
		// 不存在则创建 logs 目录
		os.MkdirAll(config.Dir.JsonData, 0o777)
	}

	// 创建一个log
	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Sys",
	})

	InstLog = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Inst",
	})

	KdataLog = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Kdata",
	})

	// 设定清除log
	mLog.Clear(mLog.ClearParam{
		Path:      config.Dir.Log,
		ClearTime: mTime.UnixTimeInt64.Day * 10,
	})

	// 将方法重载到 config 内部,便于执行
	config.LogErr = LogErr
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误 : %+v", sum)
	Log.Println(str)
}
