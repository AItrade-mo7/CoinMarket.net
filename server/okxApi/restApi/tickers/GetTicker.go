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
			BinanceTicker := okxInfo.BinanceTickerType{}
			for _, BinanceVal := range okxInfo.BinanceTickerList {
				if BinanceVal.InstID == val.InstID {
					BinanceTicker = BinanceVal
					break
				}
			}
			if len(BinanceTicker.InstID) > 2 {
				ticker := TickerCount(val, BinanceTicker)
				tickerList = append(tickerList, ticker)
			}
		}
	}

	VolumeList := BubbleVolume(tickerList) // 按照成交额排序之后
	tLen := len(VolumeList)
	if tLen > 10 {
		VolumeList = VolumeList[len(VolumeList)-10:] // 取出最后10个
	}
	okxInfo.TickerList = Reverse(VolumeList) // 翻转数组大的排在前面
	SortU_R24 := BubbleU_R24(VolumeList)     // 按照涨跌幅排序
	okxInfo.TickerU_R24 = Reverse(SortU_R24) // 24 小时涨跌幅
}
