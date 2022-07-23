package tickers

import (
	"io/ioutil"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
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

	var result okxInfo.ReqType
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.InstLog.Println("Ticker-err", result)
		return
	}

	setTicker(result.Data)

	go mFile.Write(Ticker_file, mStr.ToStr(resData))
}

func setTicker(data any) {
	var list []okxInfo.TickerType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	var tickerList []okxInfo.TickerType
	for _, val := range list {
		SPOT := okxInfo.SPOT_inst[val.InstID]
		if SPOT.State == "live" {
			ticker := TickerCount(val)
			tickerList = append(tickerList, ticker)
		}
	}

	okxInfo.TickerList = tickerList // 翻转数组大的排在前面
}
