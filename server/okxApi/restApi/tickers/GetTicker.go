package tickers

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

func GetTicker() {
	Ticker_file := mStr.Join(config.Dir.JsonData, "/Ticker.json")

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/tickers",
		Data: map[string]any{
			"instType": "SPOT",
		},
		Method:        "get",
		LocalJsonPath: Ticker_file,
		IsLocalJson:   config.SysEnv.RunMod == 1,
	})
	if err != nil {
		global.LogErr("tickers.GetTicker OKXTicker", err)
		return
	}

	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("tickers.GetTicker Ticker-err", result)
		return
	}

	setTicker(result.Data)

	mFile.Write(Ticker_file, mStr.ToStr(resData))
}

func setTicker(data any) {
	var list []mOKX.TypeOKXTicker
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	global.KdataLog.Println("inst.setTicker", len(list), "GetTicker")

	var tickerList []mOKX.TypeOKXTicker
	for _, val := range list {
		SPOT := okxInfo.SPOT_inst[val.InstID]
		if SPOT.State == "live" {
			tickerList = append(tickerList, val)
		}
	}

	VolumeList := VolumeSort(tickerList) // 按照成交额排序之后
	tLen := len(VolumeList)
	if tLen > 18 {
		VolumeList = VolumeList[tLen-15:] // 取出最后 15 个
	}
	okxInfo.OKXTickerList = []mOKX.TypeOKXTicker{}
	okxInfo.OKXTickerList = Reverse(VolumeList) // 翻转数组大的排在前面
}

// 成交量排序
func VolumeSort(arr []mOKX.TypeOKXTicker) []mOKX.TypeOKXTicker {
	size := len(arr)
	list := make([]mOKX.TypeOKXTicker, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].VolCcy24H
			b := list[j].VolCcy24H
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
func Reverse(arr []mOKX.TypeOKXTicker) []mOKX.TypeOKXTicker {
	list := make(
		[]mOKX.TypeOKXTicker,
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
