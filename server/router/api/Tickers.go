package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/gofiber/fiber/v2"
)

type TickerResType struct {
	List        []mOKX.TypeTicker                `json:"List"`        // 列表
	AnalyWhole  []mOKX.TypeWholeTickerAnaly      `json:"AnalyWhole"`  // 大盘分析结果
	AnalySingle map[string][]mOKX.AnalySliceType `json:"AnalySingle"` // 单个币种分析结果
	Unit        string                           `json:"Unit"`
	WholeDir    int                              `json:"WholeDir"`
}

func Tickers(c *fiber.Ctx) error {
	TickerRes := TickerResType{}
	TickerRes.List = okxInfo.TickerList

	return c.JSON(result.Succeed.WithData("接口开发中"))
}
