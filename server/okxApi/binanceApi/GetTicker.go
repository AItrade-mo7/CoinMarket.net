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

	var result []okxInfo.BinanceTickerType
	err = jsoniter.Unmarshal(resData, &result)
	if err != nil {
		global.InstLog.Println("BinanceTicker-err", result)
		return
	}

	SetInstID(result)

	go mFile.Write(Ticker_file, mStr.ToStr(resData))
}

func SetInstID(data []okxInfo.BinanceTickerType) {
	var list []okxInfo.BinanceTickerType
	for _, val := range data {
		Inst := val
		find := strings.Contains(val.Symbol, config.Unit)
		if find {
			Inst.InstID = strings.Replace(val.Symbol, config.Unit, config.SPOT_suffix, -1)
			SPOT := okxInfo.SPOT_inst[Inst.InstID]
			if SPOT.State == "live" {
				list = append(list, Inst)
			}
		}
	}
	okxInfo.BinanceTickerList = list
}
