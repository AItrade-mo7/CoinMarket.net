package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/gofiber/fiber/v2"
)

func GetInstAll(c *fiber.Ctx) error {
	InstAll := make(map[string]mOKX.TypeInst)
	for k, v := range okxInfo.SPOT_inst {
		InstAll[k] = v
	}
	for k, v := range okxInfo.SWAP_inst {
		InstAll[k] = v
	}
	return c.JSON(result.Succeed.WithData(InstAll))
}
