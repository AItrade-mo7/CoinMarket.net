package inst

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
)

// 获取可交易合约列表
func SWAP() {
	// 获取可交易现货列表
	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path:   "/api/v5/public/instruments",
		Method: "get",
		Data: map[string]any{
			"instType": "SWAP",
		},
	})
	if err != nil {
		global.InstLog.Fatalln("SWAP", err)
		return
	}

	// 写入日志文件
	mFile.Write(config.Dir.Log+"/SWAP.json", mStr.ToStr(resData))
}
