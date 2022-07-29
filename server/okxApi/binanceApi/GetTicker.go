package binanceApi

import (
	"io/ioutil"
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
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
		find := strings.Contains(val.Symbol, config.Unit)
		if find {
			val.InstID = strings.Replace(val.Symbol, config.Unit, config.SPOT_suffix, -1)
			SPOT := okxInfo.SPOT_inst[val.InstID]
			if SPOT.State == "live" {
				list = append(list, val)
			}
		}
	}

	VolumeList := VolumeSort(list)

	tLen := len(VolumeList)
	if tLen > 15 {
		VolumeList = VolumeList[tLen-15:] // 取出最后 15 个
	}

	okxInfo.BinanceTickerList = Reverse(VolumeList) // 翻转数组大的排在前面
}

// 成交量排序
func VolumeSort(arr []okxInfo.BinanceTickerType) []okxInfo.BinanceTickerType {
	size := len(arr)
	list := make([]okxInfo.BinanceTickerType, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].QuoteVolume
			b := list[j].QuoteVolume
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return list
}

// 翻转数组
func Reverse(arr []okxInfo.BinanceTickerType) []okxInfo.BinanceTickerType {
	list := make(
		[]okxInfo.BinanceTickerType,
		len(arr),
		len(arr)*2,
	)

	j := 0
	for i := len(arr) - 1; i > -1; i-- {
		list[j] = arr[i]
		j++
	}

	return list
}
