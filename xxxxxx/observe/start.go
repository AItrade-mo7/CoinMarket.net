package observe

import (
	"time"

	"CoinMarket.net/okxApi/instruments"
	"CoinMarket.net/okxApi/tickersAll"
	"CoinMarket.net/okxInfo"
	"CoinMarket.net/utils/hotList"
	"github.com/EasyGolang/goTools/mCycle"
)

func GetInst() {
	instruments.Start("SWAP")
	instruments.Start("SPOT")
}

func GetHotList() {
	tickersAll.Start()
}

func StoreHotList() {
	hotList.DBWrite(hotList.ListSum{
		Amount24Hot: okxInfo.Amount24Hot,
		U_R24Hot:    okxInfo.U_R24Hot,
		U_R24AbsHot: okxInfo.U_R24AbsHot,
	})
}

func Start() {
	// 请求产品数据

	mCycle.New(mCycle.Opt{
		Func:      GetInst,
		SleepTime: time.Hour * 4,
	}).Start()

	// 请求排行榜单
	mCycle.New(mCycle.Opt{
		Func:      GetHotList,
		SleepTime: time.Minute,
	}).Start()

	// 排行榜存储 10 分钟更新一次
	mCycle.New(mCycle.Opt{
		Func:      StoreHotList,
		SleepTime: time.Minute * 10,
	}).Start()
}
