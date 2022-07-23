package inst

import (
	"io/ioutil"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
)

// 获取可交易合约列表
func SWAP() {
	SWAP_file := mStr.Join(config.Dir.Log, "/SWAP.json")

	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path:   "/api/v5/public/instruments",
		Method: "get",
		Data: map[string]any{
			"instType": "SWAP",
		},
	})
	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(SWAP_file)
	}
	if err != nil {
		global.InstLog.Println("SWAP", err)
		return
	}

	// 写入日志文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}
