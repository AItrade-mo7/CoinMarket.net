package api

import (
	"CoinMarket.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func GetAnalyHistory(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithData("获取历史数据"))
}
