package ready

import (
	"os/exec"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTime"
)

func Start() {
	// 数据榜单并进行数据库存储
	SetTickerAnaly() // 默认执行一次
	go mClock.New(mClock.OptType{
		Func: SetTickerAnaly,
		Spec: "5 0,5,10,15,20,25,30,35,40,45,50,55 * * * ? ", // 5 分的整数过 5秒
	})
}

// 获取榜单数据
func SetTickerAnaly() {
	ReStartShell()   // 在这里 清理Linux 缓存
	okxApi.SetInst() // 获取并设置交易产品信息

	global.Run.Println("========= 开始获取数据 ===========")
	binanceApi.GetTicker() // 获取 okxInfo.BinanceTickerList
	tickers.GetTicker()    // 获取 okxInfo.OKXTickerList
	SetTicker()            // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据
	SetTickerNowKdata()    // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据

	okxInfo.TickerAnaly = dbType.AnalyTickerType{}
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

func ReStartShell() {
	isShellPath := mPath.Exists(config.File.ReStartShell)
	if !isShellPath {
		global.Log.Println("未找到 ReStartShell 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.ReStartShell).Output()
	global.Log.Println("执行脚本", Succeed, err)
}
