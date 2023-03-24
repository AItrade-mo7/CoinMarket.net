package binanceApi

import (
	"CoinMarket.net/server/global"
	"github.com/EasyGolang/goTools/mOKX/binance"
	jsoniter "github.com/json-iterator/go"
)

type SymbolType struct {
	Symbol                     string
	Status                     string
	BaseAsset                  string
	BaseAssetPrecision         int
	QuoteAsset                 string
	QuotePrecision             int
	QuoteAssetPrecision        int
	BaseCommissionPrecision    int
	QuoteCommissionPrecision   int
	OrderTypes                 []string
	IcebergAllowed             bool
	OcoAllowed                 bool
	QuoteOrderQtyMarketAllowed bool
	AllowTrailingStop          bool
	CancelReplaceAllowed       bool
	IsSpotTradingAllowed       bool
	IsMarginTradingAllowed     bool
	Filters                    []struct {
		FilterType            string
		MinPrice              string
		MaxPrice              string
		TickSize              string
		MultiplierUp          string
		MultiplierDown        string
		AvgPriceMins          int
		MinQty                string
		MaxQty                string
		StepSize              string
		MinNotional           string
		ApplyToMarket         bool
		Limit                 int
		MinTrailingAboveDelta int
		MaxTrailingAboveDelta int
		MinTrailingBelowDelta int
		MaxTrailingBelowDelta int
		MaxNumOrders          int
		MaxNumAlgoOrders      int
	}
	Permissions []string
}

type InstType struct {
	Timezone   string
	ServerTime int64
	RateLimits []struct {
		RateLimitType string
		Interval      string
		IntervalNum   int
		Limit         int
	}
	ExchangeFilters []interface{}
	Symbols         []SymbolType
}

func GetInst() (InstList []SymbolType) {
	resData, err := binance.FetchBinancePublic(binance.FetchBinancePublicOpt{
		Path:   "/api/v3/exchangeInfo",
		Method: "get",
	})
	if err != nil {
		global.LogErr("binanceApi.GetInst Err", err)
	}
	var result InstType
	jsoniter.Unmarshal(resData, &result)

	InstList = []SymbolType{}
	for _, val := range result.Symbols {
		if val.QuoteAsset == "USDT" && val.Status == "TRADING" {
			InstList = append(InstList, val)
		}
	}

	if len(InstList) < 10 {
		global.LogErr("binanceApi.GetInst 长度不足", len(InstList))
	}

	return
}
