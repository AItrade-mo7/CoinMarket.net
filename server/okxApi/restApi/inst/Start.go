package inst

import (
	"strings"

	"CoinMarket.net/server/global"
	"github.com/EasyGolang/goTools/mOKX"
)

var (
	SPOT_list = make(map[string]mOKX.TypeInst)
	SWAP_list = make(map[string]mOKX.TypeInst)
)

func GetInst() (InstList map[string]mOKX.TypeInst) {
	// 在这里清空数据
	InstList = make(map[string]mOKX.TypeInst)
	SPOT_inst := make(map[string]mOKX.TypeInst)
	SWAP_inst := make(map[string]mOKX.TypeInst)

	SWAP()
	SPOT()

	for key, val := range SWAP_list {
		SPOT_key := strings.Replace(key, "-SWAP", "", -1)
		SPOT := SPOT_list[SPOT_key]
		if len(SPOT.InstID) > 2 {
			InstList[SPOT.InstID] = SPOT
			InstList[val.InstID] = val
		}
	}

	global.KdataLog.Println("ready.Start inst.Start", len(SPOT_inst), len(SWAP_inst))

	return
}
