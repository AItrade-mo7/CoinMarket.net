package ready

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func SetTickerNowKdata() {
	TickerKdata := make(map[string][]mOKX.TypeKd)
	TickerList := []mOKX.TypeTicker{}

	// 一个列表最多 15 条， 一条请求结束才会是下一条
	for _, item := range okxInfo.TickerVol {
		List := mOKX.GetKdata(mOKX.GetKdataOpt{
			InstID: item.InstID,
		})

		if len(List) == config.KdataLen {
			TickerList = append(TickerList, item)
			TickerKdata[item.InstID] = List
		} else {
			global.LogErr("ready.SetTickerNowKdata Kdata 长度不足", item.InstID, len(List))
		}

	}

	okxInfo.TickerKdata = make(map[string][]mOKX.TypeKd)
	okxInfo.TickerKdata = TickerKdata

	okxInfo.TickerVol = []mOKX.TypeTicker{}
	okxInfo.TickerVol = TickerList
}
