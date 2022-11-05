package dbTask

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

func FormatKdata() {
	okxApi.SetInst() // 获取并设置交易产品信息

	GetKdata("BTC")
}

var Page = 60

func GetKdata(CcyName string) {
	InstID := mStr.Join(CcyName, "-USDT")

	for i := 0; i < 60; i++ {
		time.Sleep(time.Second / 3)
		List := okxApi.GetKdata(okxApi.GetKdataOpt{
			InstID:  InstID,
			Current: 2, // 当前页码 0 为
			After:   mTime.GetUnixInt64(),
			Size:    100,
		})

		global.Log.Println(List[0].TimeStr, List[len(List)-1].TimeStr)
	}
}
