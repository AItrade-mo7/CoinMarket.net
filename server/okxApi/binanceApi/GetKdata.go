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
	jsoniter "github.com/json-iterator/go"
)

var KdataList []mOKX.TypeKd

func GetKdata(Symbol string) {
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
		return
	}

	FormatKdata(resData, Symbol)

	go mFile.Write(Kdata_file, mStr.ToStr(resData))
}

func FormatKdata(data []byte, Symbol string) {
	var listStr [][12]string
	jsoniter.Unmarshal(data, &listStr)

	InstID := Symbol

	for _, item := range okxInfo.TickerList {
		if item.Symbol == Symbol {
			InstID = item.InstID
			break
		}
	}

	for i := len(listStr) - 1; i >= 0; i-- {
		strItem := listStr[i]
		// intItem := listInt[i]

		fmt.Println(strItem)

		kdata := mOKX.TypeKd{
			InstID: InstID,
			// Time:     mTime.MsToTime(intItem[0], "0"),
			// TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(intItem[0], "0")),
			// O:        strItem[1],
			// H:        strItem[2],
			// L:        strItem[3],
			// C:        strItem[4],
			// Vol:      strItem[5],
			// VolCcy:   strItem[7],
			Type: "BinanceKdata",
		}

		Storage(kdata)
	}
}

func Storage(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.AnalyNewKd(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)
	global.KdataLog.Println(mJson.JsonFormat(mJson.ToJson(new_Kdata)))
}
