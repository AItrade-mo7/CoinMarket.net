package okxApi

import (
	"fmt"

	"CoinMarket.net/server/okxApi/restApi/inst"
)

func GetInst() {
	InstList := inst.GetInst()

	for key := range InstList {
		fmt.Println(key)
	}
}
