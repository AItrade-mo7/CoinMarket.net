package inst

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mOKX"
)

var (
	SPOT_list = make(map[string]mOKX.TypeInst)
	SWAP_list = make(map[string]mOKX.TypeInst)
)

func Start() {
	// 在这里清空数据
	SPOT_list = make(map[string]mOKX.TypeInst)
	SWAP_list = make(map[string]mOKX.TypeInst)

	SWAP()
	SPOT()
	if len(SPOT_list) < 30 || len(SWAP_list) < 30 {
		// 正确
		global.LogErr("inst.Start 数据条目不正确", len(SPOT_list), len(SWAP_list))
		return
	}
	SPOT_inst := make(map[string]mOKX.TypeInst)
	SWAP_inst := make(map[string]mOKX.TypeInst)
	for key, val := range SWAP_list {
		SPOT_key := strings.Replace(key, "-SWAP", "", -1)
		SPOT := SPOT_list[SPOT_key]
		if len(SPOT.InstID) > 2 {
			SPOT_inst[SPOT.InstID] = SPOT
			SWAP_inst[val.InstID] = val
		}
	}
	okxInfo.SPOT_inst = SPOT_inst
	okxInfo.SWAP_inst = SWAP_inst

	global.KdataLog.Println("ready.Start inst.Start", len(okxInfo.SPOT_inst), len(okxInfo.SWAP_inst))
}
