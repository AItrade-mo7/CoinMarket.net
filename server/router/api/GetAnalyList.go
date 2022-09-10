package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/ready"
	"CoinMarket.net/server/router/result"
	"CoinMarket.net/server/utils/dbSearch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

func GetAnalyList(c *fiber.Ctx) error {
	var json dbSearch.FindParam
	mFiber.Parser(c, &json)

	if json.Current == 0 && json.Size == 300 {
		if json.Type == "Serve" {
			return c.JSON(result.Succeed.WithData(okxInfo.AnalyList_Serve))
		} else {
			return c.JSON(result.Succeed.WithData(okxInfo.AnalyList_Client))
		}
	}

	resData, err := ready.GetAnalyFirst300(json)
	if err != nil {
		return c.JSON(result.Fail.WithData(err))
	}

	return c.JSON(result.Succeed.WithData(resData))
}
