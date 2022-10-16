package ready

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

func Start() {
	// 获取 OKX 交易产品信息
	mCycle.New(mCycle.Opt{
		Func:      inst.Start,
		SleepTime: time.Hour * 4, // 每 4 时获取一次
	}).Start()
	global.KdataLog.Println("ready.Start inst.Start", len(okxInfo.SPOT_inst), len(okxInfo.SWAP_inst))

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
		go SetTickerAnalyDB()
		go SetCoinTickerDB()
		go SetCoinKdataDB("BTC")
		go SetCoinKdataDB("ETH")
	}
}
