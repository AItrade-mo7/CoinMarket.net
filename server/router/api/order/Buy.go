package order

import (
	"CoinMarket.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func Buy(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithData("Buy"))
}
