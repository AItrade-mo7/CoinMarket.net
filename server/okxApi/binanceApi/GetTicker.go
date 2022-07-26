package binanceApi

import (
	"io/ioutil"
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type BinanceTickerType struct {
	Symbol             string `json:"symbol"`
	InstID             string `json:"InstID"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int    `json:"firstId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

var BinanceTickerList []BinanceTickerType // 币安的Ticker 排行

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

	var result []BinanceTickerType
	err = jsoniter.Unmarshal(resData, &result)
	if err != nil {
		global.InstLog.Println("BinanceTicker-err", result)
		return
	}

	SetInstID(result)

	go mFile.Write(Ticker_file, mStr.ToStr(resData))
}

func SetInstID(data []BinanceTickerType) {
	var list []BinanceTickerType
	for _, val := range data {
		find := strings.Contains(val.Symbol, config.Unit)
		if find {
			val.InstID = strings.Replace(val.Symbol, config.Unit, config.SPOT_suffix, -1)
			SPOT := okxInfo.SPOT_inst[val.InstID]
			if SPOT.State == "live" {
				list = append(list, val)
			}
		}
	}

	BinanceTickerList = list
}
