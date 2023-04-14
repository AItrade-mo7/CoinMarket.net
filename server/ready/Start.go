package ready

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi"
	"github.com/EasyGolang/goTools/mClock"
	"github.com/EasyGolang/goTools/mTime"
)

// 这里只是榜单数据的爬取和搜集。
func Start() {
	StartEmail()
	// 系统重启
	go mClock.New(mClock.OptType{
		Func: SysReStart,
		Spec: "0 18 3 12,26 * ? ", // 每个月的 12 日、26 日 凌晨 5:18 重启一次 Linux 系统
	})
	// 内存清理
	go mClock.New(mClock.OptType{
		Func: ReClearShell,
		Spec: "0 18 6 * * ? ", // 每天凌晨 6:18 , 清理一次内存
	})

	// 数据榜单并进行数据库存储
	SetTickerAnaly() // 默认执行一次
	go mClock.New(mClock.OptType{
		Func: SetTickerAnaly,
		Spec: "1 1,6,11,16,21,26,31,36,41,46,51,56 * * * ? ", // 每隔5分钟比标准时间晚一分钟 过1秒 执行一次查询
	})
}

// 获取榜单数据
func SetTickerAnaly() {
	okxApi.SetInst() // 获取并设置交易产品信息

	global.Run.Println("========= 开始获取数据 ===========")

	okxApi.SetTicker() // 计算并设置综合榜单 产出 okxInfo.TickerVol 数据

	SetTickerNowKdata() // 产出 okxInfo.TickerVol 和 okxInfo.TickerKdata 数据

	// okxInfo.TickerAnaly = dbType.AnalyTickerType{} // 这里只是数据的搜集与格式化
	// okxInfo.TickerAnaly = dbType.GetAnalyTicker(tickerAnaly.TickerAnalyParam{
	// 	TickerVol:   okxInfo.TickerVol,
	// 	TickerKdata: okxInfo.TickerKdata,
	// })
	// mFile.Write(config.Dir.JsonData+"/TickerAnaly.json", mJson.ToStr(okxInfo.TickerAnaly))

	if IsOKXDataTimeScale(mTime.GetUnixInt64()) {
		go SetTickerAnalyDB()
		go SetCoinTickerDB()
		go SetCoinKdataDB("BTC")
		go SetCoinKdataDB("ETH")
	}
}
