package ready

import (
	"os/exec"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTime"
)

func Start() {
	// 数据榜单并进行数据库存储
	SetTickerAnaly()
	go mClock.New(mClock.OptType{
		Func: SetTickerAnaly,
		Spec: "5 0,5,10,15,20,25,30,35,40,45,50,55 * * * ? ", // 5 分的整数过 5秒
	})
}

// 获取榜单数据
func SetTickerAnaly() {
	if IsMongoDBTimeScale(mTime.GetUnixInt64()) {
		ReStartMongoDB() // 在这里重启数据库
	}

	inst.Start() // 获取交易产品信息

	global.Run.Println("========= 开始获取数据 ===========")
	binanceApi.GetTicker() // 获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据
	SetTickerKdata()       // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据

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

func ReStartMongoDB() {
	isShellPath := mPath.Exists(config.File.ReStartShell)
	if !isShellPath {
		global.Log.Println("未找到 ReStartShell 脚本")
		return
	}

	Succeed, err := exec.Command("/bin/bash", config.File.ReStartShell).Output()
	global.Log.Println("执行脚本", Succeed, err)
}
