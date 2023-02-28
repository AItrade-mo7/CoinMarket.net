package global

import (
	"fmt"
	"log"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/utils/taskPush"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mLog"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

var (
	Log             *log.Logger // 系统日志& 重大错误或者事件
	Run             *log.Logger //  运行日志
	KdataLog        *log.Logger //  OKX Kdata 日志
	BinanceKdataLog *log.Logger //  币安 Kdata 日志
)

func LogInit() {
	// 创建一个log
	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Sys",
	})

	Run = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Run",
	})

	KdataLog = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Kdata",
	})

	BinanceKdataLog = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "BinanceKdata",
	})

	// 设定清除log
	mLog.Clear(mLog.ClearParam{
		Path:      config.Dir.Log,
		ClearTime: mTime.UnixTimeInt64.Day * 10,
	})

	// 将方法重载到 config 内部,便于执行
	config.LogErr = LogErr
	config.Log = Log
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误: %+v", sum)
	Log.Println(str)

	message := ""
	if len(sum) > 0 {
		message = mStr.ToStr(sum[0])
	}
	content := mJson.Format(sum)

	err := taskPush.SysEmail(taskPush.SysEmailOpt{
		From:        config.SysName,
		Subject:     "系统错误",
		Title:       config.SysName + " 系统出错",
		Message:     message,
		Content:     content,
		Description: "出现系统错误",
	})
	Log.Println("邮件已发送", err)
}
