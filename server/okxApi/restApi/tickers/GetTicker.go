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

type OKXTickerType struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
}

var OKXTickerList []OKXTickerType // okx的Ticker

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
	var list []OKXTickerType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	var tickerList []OKXTickerType
	for _, val := range list {
		SPOT := okxInfo.SPOT_inst[val.InstID]
		if SPOT.State == "live" {
			tickerList = append(tickerList, val)
		}
	}

	VolumeList := VolumeSort(tickerList) // 按照成交额排序之后
	tLen := len(VolumeList)
	if tLen > 10 {
		VolumeList = VolumeList[len(VolumeList)-10:] // 取出最后10个
	}
	OKXTickerList = Reverse(VolumeList) // 翻转数组大的排在前面
}
