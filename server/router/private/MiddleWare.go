package private

import (
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/middle"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", config.SysName+"/api/private")

	// Token 验证
	_, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	return c.Next()
}
