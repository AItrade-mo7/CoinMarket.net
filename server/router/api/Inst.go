package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type InstParam struct {
	InstType string `json:"instType"`
}

func Inst(c *fiber.Ctx) error {
	var json InstParam
	mFiber.Parser(c, &json)

	if json.InstType == "SPOT" {
		return c.JSON(result.Succeed.WithData(okxInfo.SPOT_inst))
	}

	if json.InstType == "SWAP" {
		return c.JSON(result.Succeed.WithData(okxInfo.SWAP_inst))
	}

	return c.JSON(result.Fail.WithMsg("缺少 InstType"))
}
