package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func GetInstAll(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithData(okxInfo.Inst))
}
