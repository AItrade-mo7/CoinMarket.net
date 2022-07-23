package inst

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi/okxInfo"
)

var (
	SPOT_list = make(map[string]okxInfo.InstType)
	SWAP_list = make(map[string]okxInfo.InstType)
)

func Start() {
	SWAP()
	SPOT()

	if len(SPOT_list) < 30 || len(SWAP_list) < 30 {
		// 正确
		global.InstLog.Println("数据条目不正确", len(SPOT_list), len(SWAP_list))
		return
	}

	for key, val := range SWAP_list {
		fmt.Println(key, val.InstType)
	}
}
