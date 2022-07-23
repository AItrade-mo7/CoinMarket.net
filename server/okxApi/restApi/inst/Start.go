package inst

import "CoinMarket.net/server/okxApi/okxInfo"

var (
	SPOT_list []okxInfo.InstType
	SWAP_list []okxInfo.InstType
)

func Start() {
	SWAP()
	SPOT()
}
