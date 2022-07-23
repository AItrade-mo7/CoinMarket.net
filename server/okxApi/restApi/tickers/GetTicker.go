package tickers

import (
	"io/ioutil"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
)

func GetTicker() {
	Ticker_file := mStr.Join(config.Dir.JsonData, "/Ticker.json")

	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path:   "/api/v5/market/tickers",
		Method: "get",
		Data: map[string]any{
			"instType": "SPOT",
		},
	})
	// 本地模式
	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(Ticker_file)
	}

	if err != nil {
		global.InstLog.Println("OKXTicker", err)
		return
	}

	go mFile.Write(Ticker_file, mStr.ToStr(resData))
}
