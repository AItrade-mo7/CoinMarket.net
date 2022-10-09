package api

import (
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

func GetAnalyList(c *fiber.Ctx) error {
	var json dbSearch.FindParam
	mFiber.Parser(c, &json)

	return c.JSON(result.Succeed.WithData("接口开发中"))
}
