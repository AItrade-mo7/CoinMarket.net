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
	SWAP()
	SPOT()
	if len(SPOT_list) < 30 || len(SWAP_list) < 30 {
		// 正确
		global.LogErr("inst 数据条目不正确", len(SPOT_list), len(SWAP_list))
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
}
