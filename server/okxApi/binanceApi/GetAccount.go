package binanceApi

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mBinance"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type TypeAccount struct {
	FeeTier                     int    `json:"feeTier"`
	CanTrade                    bool   `json:"canTrade"`
	CanDeposit                  bool   `json:"canDeposit"`
	CanWithdraw                 bool   `json:"canWithdraw"`
	UpdateTime                  int    `json:"updateTime"`
	MultiAssetsMargin           bool   `json:"multiAssetsMargin"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintMargin            string `json:"totalMaintMargin"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string `json:"totalCrossUnPnl"`
	AvailableBalance            string `json:"availableBalance"`
	MaxWithdrawAmount           string `json:"maxWithdrawAmount"`
	Assets                      []struct {
		Asset                  string `json:"asset"`
		WalletBalance          string `json:"walletBalance"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		MarginBalance          string `json:"marginBalance"`
		MaintMargin            string `json:"maintMargin"`
		InitialMargin          string `json:"initialMargin"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
		CrossWalletBalance     string `json:"crossWalletBalance"`
		CrossUnPnl             string `json:"crossUnPnl"`
		AvailableBalance       string `json:"availableBalance"`
		MarginAvailable        bool   `json:"marginAvailable"`
		UpdateTime             int    `json:"updateTime"`
	} `json:"assets"`
	Positions []struct {
		Symbol                 string `json:"symbol"`
		InitialMargin          string `json:"initialMargin"`
		MaintMargin            string `json:"maintMargin"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		Leverage               string `json:"leverage"`
		Isolated               bool   `json:"isolated"`
		EntryPrice             string `json:"entryPrice"`
		MaxNotional            string `json:"maxNotional"`
		PositionSide           string `json:"positionSide"`
		PositionAmt            string `json:"positionAmt"`
		Notional               string `json:"notional"`
		IsolatedWallet         string `json:"isolatedWallet"`
		UpdateTime             int    `json:"updateTime"`
		BidNotional            string `json:"bidNotional"`
		AskNotional            string `json:"askNotional"`
	} `json:"positions"`
}

func GetAccount() (resultData mBinance.PositionType) {
	Kdata_file := mStr.Join(config.Dir.JsonData, "/阔盈-Account.json")

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

	var result TypeAccount
	jsoniter.Unmarshal(resData, &result)

	PositionAmt := 0
	PositionSymbol := ""
	Profit := ""

	for _, val := range result.Positions {
		if mCount.Le(val.PositionAmt, "0") != 0 {
			PositionAmt = mCount.Le(val.PositionAmt, "0")
			PositionSymbol = val.Symbol
			Profit = val.UnrealizedProfit
		}
	}

	PositionInst := okxInfo.Inst[PositionSymbol].BaseCcy
	if len(PositionInst) > 1 {
		PositionInst = PositionInst + config.SWAP_suffix
	}

	resultData = mBinance.PositionType{
		InstID: PositionInst,
		Dir:    PositionAmt,
		Profit: Profit,
	}

	resultData.CreateTime = mTime.GetUnixInt64()
	resultData.CreateTimeStr = mTime.UnixFormat(resultData.CreateTime)

	mFile.Write(Kdata_file, mStr.ToStr(resData))

	return
}
