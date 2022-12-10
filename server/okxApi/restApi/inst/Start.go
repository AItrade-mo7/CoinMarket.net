package inst

import (
	"strings"

	"github.com/EasyGolang/goTools/mOKX"
)

var (
	SPOT_list = make(map[string]mOKX.TypeInst)
	SWAP_list = make(map[string]mOKX.TypeInst)
)

func GetInst() (InstList map[string]mOKX.TypeInst) {
	// 在这里清空数据
	InstList = make(map[string]mOKX.TypeInst)
	SPOT_list = make(map[string]mOKX.TypeInst)
	SWAP_list = make(map[string]mOKX.TypeInst)

	SWAP()
	SPOT()

	// 清洗一遍，只留下合约版本
	for key, val := range SWAP_list {
		SPOT_key := strings.Replace(key, "-SWAP", "", -1)
		SPOT := SPOT_list[SPOT_key]
		if len(SPOT.InstID) > 2 {
			InstList[SPOT.InstID] = SPOT
			InstList[val.InstID] = val
		}
	}

	return
}
