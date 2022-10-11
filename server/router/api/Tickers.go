package api

import (
	"CoinMarket.net/server/global/apiType"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/tickerAnaly"
	"github.com/gofiber/fiber/v2"
)

func Tickers(c *fiber.Ctx) error {
	TickerRes := apiType.AnalyTickerType{}

	AnalyResult := tickerAnaly.GetAnaly(tickerAnaly.TickerAnalyParam{
		TickerList:  okxInfo.TickerList,
		MarketKdata: okxInfo.TickerKdata,
	})

	TickerRes.List = okxInfo.TickerList
	TickerRes.AnalyWhole = AnalyResult.AnalyWhole
	TickerRes.AnalySingle = AnalyResult.AnalySingle
	TickerRes.WholeDir = AnalyResult.WholeDir
	TickerRes.Unit = config.Unit

	return c.JSON(result.Succeed.WithData(TickerRes))
}
