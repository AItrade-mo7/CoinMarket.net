package global

import (
	"fmt"
	"log"
	"os"
	"time"

	"CoinMarket.net/config"
	"CoinMarket.net/tmpl"
	"github.com/EasyGolang/goTools/mLog"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTime"
)

var (
	Log        *log.Logger // 系统日志& 重大错误或者事件
	LogWss     *log.Logger // WSS 监听记录 - 不间断的会有
	LogRestApi *log.Logger // ResetAPI 的请求与数据记录
	LogInst    *log.Logger // 产品数据的请求与记录
	LogHotList *log.Logger // 行情数据
)

func LogInt() {
	// 检测 logs 目录
	isLogPath := mPath.Exists(config.LogPath)
	if !isLogPath {
		// 不存在则创建 logs 目录
		os.Mkdir(config.LogPath, 0o777)
	}

	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.LogPath,
		Name: "Sys",
	})

	LogWss = mLog.NewLog(mLog.NewLogParam{
		Path: config.LogPath,
		Name: "Wss",
	})

	LogRestApi = mLog.NewLog(mLog.NewLogParam{
		Path: config.LogPath,
		Name: "RestApi",
	})

	LogInst = mLog.NewLog(mLog.NewLogParam{
		Path: config.LogPath,
		Name: "Inst",
	})

	LogHotList = mLog.NewLog(mLog.NewLogParam{
		Path: config.LogPath,
		Name: "HotList",
	})

	mLog.Clear(mLog.ClearParam{
		Path:      config.LogPath,
		ClearTime: mTime.UnixTimeInt64.Day * 10,
	})
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误 : %+v", sum)
	Log.Println(str)

	Email := Email(EmailOpt{
		To:       UserEnv.Email.To,
		Subject:  "LogErr",
		Template: tmpl.Email,
		SendData: struct {
			Message string
			SysTime time.Time
			RunMod  int
		}{
			Message: str,
			SysTime: time.Now(),
			RunMod:  ServerEnv.RunMod,
		},
	})

	go Email.Send()
}
