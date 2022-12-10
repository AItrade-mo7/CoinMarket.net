package okxApi

import (
	"CoinMarket.net/server/okxApi/restApi/inst"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

func SetInst() {
	InstList := inst.GetInst()

	MergeInstList := make(map[string]mOKX.TypeInst)

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

	okxInfo.Inst = make(map[string]mOKX.TypeInst) // 清理产品信息
	okxInfo.Inst = MergeInstList                  // 获取并设置交易产品信息
}
