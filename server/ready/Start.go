package ready

import (
	"log"
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
)

func Start() {
	// 获取 OKX 交易产品信息
	mCycle.New(mCycle.Opt{
		Func:      inst.Start,
		SleepTime: time.Hour * 4, // 每 4 时获取一次
	}).Start()
	global.KdataLog.Println("ready.Start inst.Start", len(okxInfo.SPOT_inst), len(okxInfo.SWAP_inst))

	// 数据榜单并进行数据库存储
	go mClock.New(mClock.OptType{
		Func: SetDBTickerData,
		Spec: "0,5,10,15,20,25,30,35,40,45,50,55 * * * * ? ",
	})
}

// 获取历史数据并存储
func SetDBTickerData() {
	log.Println("开始获取数据")

	// time.Sleep(time.Second / 3)
	SetTickerAnaly() //  产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 以及 okxInfo.TickerAnaly

	nowTime := mOKX.GetKdataTime(okxInfo.TickerKdata)

	log.Println("数据获取结束", nowTime)

	// go SetTickerAnalyDB()
	// go SetCoinTickerDB()
	// go SetCoinKdataDB("BTC")
	// go SetCoinKdataDB("ETH")
}

// 获取榜单数据
func SetTickerAnaly() {
	binanceApi.GetTicker() // 获取币安的 Ticker
	tickers.GetTicker()    // 获取 okx 的Ticker
	SetTicker()            // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据
	SetTickerKdata()       // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据
	global.Run.Println(
		"========= 开始分析 ===========",
		len(okxInfo.TickerVol),
		len(okxInfo.TickerKdata),
		len(okxInfo.TickerKdata["BTC-USDT"]),
	)
	okxInfo.TickerAnaly = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
		TickerVol:   okxInfo.TickerVol,
		TickerKdata: okxInfo.TickerKdata,
	})
}
