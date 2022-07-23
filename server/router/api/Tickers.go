package api

import (
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/gofiber/fiber/v2"
)

type TickersParam struct {
	SortType string `json:"SortType"`
}

func Tickers(c *fiber.Ctx) error {
	var json TickersParam
	mFiber.Parser(c, &json)

	if json.SortType == "U_R24" {
		return c.JSON(result.Succeed.WithData(okxInfo.TickerU_R24))
	}

	if json.SortType == "Amount" {
		return c.JSON(result.Succeed.WithData(okxInfo.TickerList))
	}

	return c.JSON(result.Fail.WithMsg("缺少 SortType"))
}
