package api

import (
	"CoinMarket.net/server/ready"
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

func GetAnalyList(c *fiber.Ctx) error {
	var json dbSearch.FindParam
	mFiber.Parser(c, &json)

	resData, err := ready.GetAnalyFirst300(json)
	if err != nil {
		return c.JSON(result.Fail.WithData(err))
	}

	return c.JSON(result.Succeed.WithData(resData))
}
