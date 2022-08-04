package binanceApi

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

func GetKdata(Symbol string) {
	Kdata_file := mStr.Join(config.Dir.JsonData, "/B-", Symbol, ".json")
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

	go mFile.Write(Kdata_file, mStr.ToStr(resData))
}
