package ready

import (
	"fmt"

	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/kdata"
)

func TickerKData() {
	MaxNum := 4 // 去除前 6 条数据
	for key, item := range okxInfo.TickerList {
		fmt.Println(key, item.InstID)
		if key > MaxNum {
			break
		}
	}

	kdata.GetKdata(okxInfo.TickerList[0].InstID)
}
