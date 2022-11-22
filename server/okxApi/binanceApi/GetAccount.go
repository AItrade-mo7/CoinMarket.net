package binanceApi

import (
	"context"
	"fmt"

	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/adshao/go-binance/v2"
)

func GetAccount() {
	var (
		apiKey    = config.BinanceKey.ApiKey
		secretKey = config.BinanceKey.SecretKey
	)
	client := binance.NewClient(apiKey, secretKey)

	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	mFile.Write(config.Dir.JsonData+"/Account.json", mJson.ToStr(res))
}
