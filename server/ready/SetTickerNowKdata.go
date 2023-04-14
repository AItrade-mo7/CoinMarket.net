package ready

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

var Page = 3 // 4 页的数据

func SetTickerNowKdata() {
	TickerKdata := make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}

	// 一个列表最多 15 条， 一条请求结束才会是下一条
	for _, item := range okxInfo.TickerVol {
		List := Get400Kdata(item.InstID)
		if len(List) == config.KdataLen*(Page+1) {
			TickerList = append(TickerList, item)
			TickerKdata[item.InstID] = List
		} else {
			global.LogErr("ready.SetTickerNowKdata Kdata 长度不足，从数据中除名", item.InstID, len(List))
		}
	}

	okxInfo.TickerKdata = make(map[string][]mOKX.TypeKd)
	okxInfo.TickerKdata = TickerKdata

	okxInfo.TickerVol = []mOKX.TypeTicker{}
	okxInfo.TickerVol = TickerList
}

func Get400Kdata(InstID string) []mOKX.TypeKd {
	AllList := []mOKX.TypeKd{}
	for i := Page; i >= 0; i-- {
		List := mOKX.GetKdata(mOKX.GetKdataOpt{
			InstID: InstID,
			Page:   i,
			After:  mTime.GetUnixInt64(),
		})
		AllList = append(AllList, List...)
	}

	for key, val := range AllList {
		preIndex := key - 1
		if preIndex < 0 {
			preIndex = 0
		}
		preItem := AllList[preIndex]
		nowItem := AllList[key]
		if key > 0 {
			if nowItem.TimeUnix-preItem.TimeUnix != mTime.UnixTimeInt64.Hour {
				global.LogErr("数据检查出错 backTest.CheckKdataList", val.InstID, val.TimeStr, key)
				AllList = []mOKX.TypeKd{}
				break
			}
		}
	}
	time.Sleep(time.Second / 2)

	return AllList
}
