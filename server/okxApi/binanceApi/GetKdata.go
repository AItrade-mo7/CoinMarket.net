package binanceApi

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

var KdataList []mOKX.TypeKd

func GetKdata(Symbol string) []mOKX.TypeKd {
	Kdata_file := mStr.Join(config.Dir.JsonData, "/B-", Symbol, ".json")

	KdataList = []mOKX.TypeKd{}
	resData, err := mOKX.FetchBinance(mOKX.FetchBinanceOpt{
		Path:   "/api/v3/klines",
		Method: "get",
		Data: map[string]any{
			"symbol":   Symbol,
			"interval": "15m",
			"limit":    300,
		},
		LocalJsonPath: Kdata_file,
		IsLocalJson:   config.AppEnv.RunMod == 1,
	})
	if err != nil {
		global.InstLog.Println("BinanceTicker", err)
		return nil
	}

	FormatKdata(resData, Symbol)

	go mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
}

func FormatKdata(data []byte, Symbol string) {
	var listStr [][12]any
	jsoniter.Unmarshal(data, &listStr)

	InstID := Symbol

	for _, item := range okxInfo.TickerList {
		if item.Symbol == Symbol {
			InstID = item.InstID
			break
		}
	}

	for _, item := range listStr {
		TimeStr := mStr.ToStr(mJson.ToJson(item[0]))

		kdata := mOKX.TypeKd{
			InstID:   InstID,
			Time:     mTime.MsToTime(TimeStr, "0"),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(TimeStr, "0")),
			O:        mStr.ToStr(item[1]),
			H:        mStr.ToStr(item[2]),
			L:        mStr.ToStr(item[3]),
			C:        mStr.ToStr(item[4]),
			Vol:      mStr.ToStr(item[5]),
			VolCcy:   mStr.ToStr(item[7]),
			Type:     "BinanceKdata",
		}

		fmt.Println(kdata.Time)

		Storage(kdata)
	}
}

func Storage(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.AnalyNewKd(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)

	global.BinanceKdataLog.Println(mJson.JsonFormat(mJson.ToJson(new_Kdata)))
}