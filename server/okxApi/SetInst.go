package okxApi

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

func SetInst() {
	InstList := inst.GetInst()

	binanceInstList := binanceApi.GetInst()

	if len(InstList) < 10 || len(binanceInstList) < 10 {
		global.LogErr("okxApi.SetInst 数量不足1", len(InstList), len(binanceInstList))
		return
	}

	MergeInstList := make(map[string]mOKX.TypeInst)
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

	// 给合约 添加 Symbol
	for _, val := range InstList {
		if val.InstType == "SWAP" {
			SPOT := MergeInstList[val.Uly]
			if len(SPOT.Symbol) > 4 {
				val.Symbol = SPOT.Symbol
				MergeInstList[val.InstID] = val
			}
		}
	}

	okxInfo.Inst = make(map[string]mOKX.TypeInst) // 清理产品信息
	if len(MergeInstList) > 10 {
		okxInfo.Inst = MergeInstList // 获取并设置交易产品信息
	} else {
		global.LogErr("okxApi.SetInst 数量不足2", len(MergeInstList))
	}
}
