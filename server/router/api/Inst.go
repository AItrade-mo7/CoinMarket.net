package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type InstParam struct {
	TypeInst string `json:"TypeInst"`
}

func Inst(c *fiber.Ctx) error {
	var json InstParam
	mFiber.Parser(c, &json)

	if json.TypeInst == "SPOT" {
		return c.JSON(result.Succeed.WithData(okxInfo.SPOT_inst))
	}

	if json.TypeInst == "SWAP" {
		return c.JSON(result.Succeed.WithData(okxInfo.SWAP_inst))
	}

	return c.JSON(result.Fail.WithMsg("缺少 TypeInst"))
}
