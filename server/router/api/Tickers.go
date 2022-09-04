package api

import (
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func Tickers(c *fiber.Ctx) error {
	TickerDB := dbType.GetTickerDB()
	return c.JSON(result.Succeed.WithData(TickerDB))
}
