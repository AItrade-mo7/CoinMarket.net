package api

import (
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/gofiber/fiber/v2"
)

type TickersParam struct {
	SortType string `json:"SortType"` //  Amount  || U_R24
}

type TickerResType struct {
	List        []mOKX.TypeTicker                `json:"List"`        // 列表
	AnalyWhole  []mOKX.TypeWholeTickerAnaly      `json:"AnalyWhole"`  // 大盘分析结果
	AnalySingle map[string][]mOKX.AnalySliceType `json:"AnalySingle"` // 单个币种分析结果
	Unit        string                           `json:"Unit"`
	WholeDir    int                              `json:"WholeDir"`
}

func Tickers(c *fiber.Ctx) error {
	var json TickersParam
	mFiber.Parser(c, &json)

	TickerRes := TickerResType{}
	TickerRes.List = okxInfo.TickerList
	TickerRes.AnalyWhole = okxInfo.TickerAnalyWhole
	TickerRes.AnalySingle = okxInfo.TickerAnalySingle
	TickerRes.Unit = config.Unit
	TickerRes.WholeDir = okxInfo.WholeDir

	if json.SortType == "U_R24" {
		TickerRes.List = okxInfo.TickerU_R24
	}

	return c.JSON(result.Succeed.WithData(TickerRes))
}
