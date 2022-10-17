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
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTime"
)

func Start() {
	// 设定数据库重启
	ReStartMongoDB()
	go mClock.New(mClock.OptType{
		Func: ReStartMongoDB,
		Spec: "0 8 3,7,11,15,19,23 * * ? ", // 数据库重启
	})

	// 获取 OKX 交易产品信息
	inst.Start()
	go mClock.New(mClock.OptType{
		Func: inst.Start,
		Spec: "0 7 4,9,16,21 * * ? ",
	})

	// 数据榜单并进行数据库存储
	SetTickerAnaly()
	go mClock.New(mClock.OptType{
		Func: SetTickerAnaly,
		Spec: "5 0,5,10,15,20,25,30,35,40,45,50,55 * * * ? ", // 5 分的整数过 5秒
	})
}

// 获取榜单数据
func SetTickerAnaly() {
	global.Run.Println("========= 开始获取数据 ===========")

	binanceApi.GetTicker() // 获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据
	SetTickerKdata()       // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据

	global.Run.Println(
		"== 开始分析 ==",
		len(okxInfo.TickerVol),
		len(okxInfo.TickerKdata),
	)

	okxInfo.TickerAnaly = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
		TickerVol:   okxInfo.TickerVol,
		TickerKdata: okxInfo.TickerKdata,
	})

	global.Run.Println(
		"== 分析结束 ==",
		mTime.UnixFormat(mOKX.GetKdataTime(okxInfo.TickerKdata)),
		len(okxInfo.TickerAnaly.TickerVol),
		len(okxInfo.TickerAnaly.AnalyWhole),
		len(okxInfo.TickerAnaly.AnalySingle),
		len(okxInfo.TickerAnaly.Unit),
		okxInfo.TickerAnaly.WholeDir,
		okxInfo.TickerAnaly.TimeID,
	)

	if IsTimeScale(mTime.GetUnixInt64()) {
		// go SetTickerAnalyDB()
		// go SetCoinTickerDB()
		// go SetCoinKdataDB("BTC")
		// go SetCoinKdataDB("ETH")
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
