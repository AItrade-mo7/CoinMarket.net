package public

import (
	"CoinMarket.net/server/router/middle"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", "DataCenter.net/CoinMarket/public")

	// 授权验证
	err := middle.EncryptAuth(c)
	if err != nil {
		return c.JSON(result.ErrAuth.WithData(mStr.ToStr(err)))
	}

	return c.Next()
}
