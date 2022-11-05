package okxApi

import (
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

func GetInst() (MergeInstList map[string]mOKX.TypeInst) {
	binanceInstList := binanceApi.GetInst()
	InstList := inst.GetInst()

	MergeInstList = make(map[string]mOKX.TypeInst)

	// 整理现货
	for _, okxItem := range InstList {
		Symbol := mStr.Join(okxItem.BaseCcy, okxItem.QuoteCcy)
		for _, binanceItem := range binanceInstList {
			if binanceItem.Symbol == Symbol {
				okxItem.Symbol = binanceItem.Symbol
				MergeInstList[okxItem.InstID] = okxItem
				break
			}
		}
	}
	// 添加合约
	for _, val := range InstList {
		if val.InstType == "SWAP" {
			SPOT := MergeInstList[val.Uly]
			val.Symbol = SPOT.Symbol
			if len(SPOT.Symbol) > 4 {
				MergeInstList[val.InstID] = val
			}
		}
	}

	return
}
