package binanceApi

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mOKX/binance"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

// 币安的 ticker 数据
func GetTicker() (TickerList []mOKX.TypeBinanceTicker) {
	Ticker_file := mStr.Join(config.Dir.JsonData, "/Bnb_Ticker.json")
	resData, err := binance.FetchBinancePublic(binance.FetchBinancePublicOpt{
		Path:   "/api/v3/ticker/24hr",
		Method: "get",
	})
	if err != nil {
		global.LogErr("binanceApi.GetTicker BinanceTicker", err)
		return
	}

	var result []mOKX.TypeBinanceTicker
	err = jsoniter.Unmarshal(resData, &result)
	if err != nil {
		global.LogErr("binanceApi.GetTicker BinanceTicker-err", result)
		return
	}

	TickerList = SetInstID(result)

	okxInfo.BinanceTickerList = []mOKX.TypeBinanceTicker{}
	okxInfo.BinanceTickerList = TickerList

	mFile.Write(Ticker_file, mStr.ToStr(resData))

	return
}

func SetInstID(data []mOKX.TypeBinanceTicker) (TickerList []mOKX.TypeBinanceTicker) {
	var list []mOKX.TypeBinanceTicker

	global.BinanceKdataLog.Println("binanceApi.SetInstID", len(data), "GetTicker")

	for _, val := range data {
		find := strings.Contains(val.Symbol, config.Unit)
		if find {
			InstID := strings.Replace(val.Symbol, "USDT", "-USDT", 1)
			SPOT := okxInfo.Inst[InstID]
			val.InstID = SPOT.InstID
			if len(SPOT.Symbol) > 3 {
				list = append(list, val)
			}
		}
	}

	VolumeList := VolumeSort(list)

	tLen := len(VolumeList)
	if tLen > 18 {
		VolumeList = VolumeList[tLen-15:] // 取出最后 15 个
	}

	TickerList = Reverse(VolumeList) // 翻转数组大的排在前面
	return
}

// 成交量排序
func VolumeSort(arr []mOKX.TypeBinanceTicker) []mOKX.TypeBinanceTicker {
	size := len(arr)
	list := make([]mOKX.TypeBinanceTicker, size)
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
func Reverse(arr []mOKX.TypeBinanceTicker) []mOKX.TypeBinanceTicker {
	list := make(
		[]mOKX.TypeBinanceTicker,
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
