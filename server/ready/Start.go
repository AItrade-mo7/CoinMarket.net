package ready

import (
	"os/exec"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnaly"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

// 这里只是榜单数据的爬取和搜集。
func Start() {
	// 发送启动邮件
	go global.Email(global.EmailOpt{
		To:       config.Email.To,
		Subject:  "ServeStart",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "服务启动",
			SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
		},
	}).Send()
	// 系统重启
	go mClock.New(mClock.OptType{
		Func: SysReStart,
		Spec: "0 18 5 3,9,15,21,27 * ? ", // 每个月的 3 日、9 日 每隔5天的凌晨 5:18 重启一次 Linux 系统
	})

	// 内存清理
	go mClock.New(mClock.OptType{
		Func: ReClearShell,
		Spec: "0 18 3 * * ? ", // 每天凌晨 3:18 ，数据库重启
	})

	// 数据榜单并进行数据库存储
	SetTickerAnaly() // 默认执行一次
	go mClock.New(mClock.OptType{
		Func: SetTickerAnaly,
		Spec: "1 1,6,11,16,21,26,31,36,41,46,51,56 * * * ? ", // 每隔5分钟比标准时间晚一分钟 过1秒 执行查询
	})
}

// 获取榜单数据
func SetTickerAnaly() {
	okxApi.SetInst() // 获取并设置交易产品信息

	global.Run.Println("========= 开始获取数据 ===========")
	go SetBinancePosition()

	okxApi.SetTicker() // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据

	SetTickerNowKdata() // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据

	okxInfo.TickerAnaly = dbType.AnalyTickerType{} // 这里只是数据的搜集与格式化
	okxInfo.TickerAnaly = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
		TickerVol:   okxInfo.TickerVol,
		TickerKdata: okxInfo.TickerKdata,
	})

	if IsOKXDataTimeScale(mTime.GetUnixInt64()) {
		go SetTickerAnalyDB()
		go SetCoinTickerDB()
		go SetCoinKdataDB("BTC")
		go SetCoinKdataDB("ETH")
	}
}

func ReClearShell() {
	isShellPath := mPath.Exists(config.File.ReClearShell)
	if !isShellPath {
		global.Log.Println("未找到 ReClearShell 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.ReClearShell).Output()
	global.Log.Println("执行脚本", Succeed, err)

	go global.Email(global.EmailOpt{
		To:       config.Email.To,
		Subject:  "数据库重启并执行清理",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: mStr.ToStr(Succeed),
			SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
		},
	}).Send()
}

func SysReStart() {
	go global.Email(global.EmailOpt{
		To:       config.Email.To,
		Subject:  "Linux系统即将重启",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: "Linux系统 10 秒钟后 将自动进行重启",
			SysTime: mTime.UnixFormat(mTime.GetUnixInt64()),
		},
	}).Send()

	time.Sleep(time.Second * 10)

	isShellPath := mPath.Exists(config.File.SysReStart)
	if !isShellPath {
		global.Log.Println("未找到 SysReStart 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.SysReStart).Output()
	global.Log.Println("执行脚本", Succeed, err)
}
