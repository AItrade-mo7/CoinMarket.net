package api

import (
	"CoinMarket.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func Tickers(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithData("接口开发中"))
}
