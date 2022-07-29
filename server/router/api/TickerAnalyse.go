package api

import (
	"CoinMarket.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func TickerAnalyse(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithData("学习"))
}
