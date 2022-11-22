package binanceApi

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTime"
)

func GetAccount() {
	Timestamp := mTime.GetUnixInt64()
	fmt.Println(Timestamp)
}
