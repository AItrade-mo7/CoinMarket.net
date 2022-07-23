package binanceApi

import (
	"io/ioutil"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
)

// 币安的 ticker 数据
func GetTicker() {
	Ticker_file := mStr.Join(config.Dir.JsonData, "/BinanceTicker.json")
	resData, err := Fetch(FetchOpt{
		Path:   "/api/v3/ticker/24hr",
		Method: "get",
	})
	// 本地模式
	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(Ticker_file)
	}
	if err != nil {
		global.InstLog.Println("BinanceTicker", err)
		return
	}

	go mFile.Write(Ticker_file, mStr.ToStr(resData))
}
