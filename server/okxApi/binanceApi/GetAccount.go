package binanceApi

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mBinance"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
)

func GetAccount() {
	Kdata_file := mStr.Join(config.Dir.JsonData, "/B-Account.json")

	resData, err := mBinance.FetchBinance(mBinance.OptFetchBinance{
		Path:   "/fapi/v2/account",
		Method: "get",
		BinanceKey: mBinance.TypeBinanceKey{
			ApiKey:    config.BinanceKey.ApiKey,
			SecretKey: config.BinanceKey.SecretKey,
		},
	})
	if err != nil {
		global.LogErr("binanceApi.GetAccount Err", err)
	}

	mFile.Write(Kdata_file, mStr.ToStr(resData))
}
