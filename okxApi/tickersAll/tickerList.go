package tickersAll

import (
	"strings"

	"CoinMarket.net/okxInfo"
	"CoinMarket.net/utils/hotList"
	"github.com/EasyGolang/goTools/mCount"
)

func SetHotList(list []okxInfo.HotInfo) {
	hList := []okxInfo.HotInfo{}

	for _, val := range list {
		ticker := TickerCount(val)
		if len(ticker.CcyName) < 1 {
			continue
		}
		hList = append(hList, ticker)
	}

	okxInfo.Amount24Hot = hotList.SortAmount(hList)
	okxInfo.U_R24Hot = hotList.SortU_R(hList)
	okxInfo.U_R24AbsHot = hotList.SortU_RAbs(hList)
}

var AmountCut = "100000000"

func TickerCount(data okxInfo.HotInfo) okxInfo.HotInfo {
	ticker := data
	// 对应币种
	ticker.CcyName = strings.Replace(ticker.InstID, "-USDT-SWAP", "", 1)
	// 涨幅
	ticker.U_R24 = mCount.RoseCent(ticker.Last, ticker.Open24H) // 24 时涨幅
	// 平均价格
	ticker.Average24 = mCount.Average([]string{
		ticker.Open24H,
		ticker.High24H,
		ticker.Low24H,
		ticker.Last,
		ticker.SodUtc0,
		ticker.SodUtc8,
	})

	ticker.Amount = mCount.Mul(
		ticker.Average24,
		ticker.VolCcy24H,
	)

	ticker.SPOT = hotList.FindSPOTInst(ticker)
	ticker.SWAP = hotList.FindSWAPInst(ticker)

	// 资金金额大于一个亿
	if mCount.Le(ticker.Amount, AmountCut) < 0 {
		return okxInfo.HotInfo{}
	}

	if len(ticker.SPOT.BaseCcy) < 1 {
		return okxInfo.HotInfo{}
	}
	if len(ticker.SWAP.CtValCcy) < 1 {
		return okxInfo.HotInfo{}
	}
	return ticker
}
