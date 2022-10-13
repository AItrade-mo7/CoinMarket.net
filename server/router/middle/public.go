package middle

import (
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

func Public(c *fiber.Ctx) error {
	// 添加访问头
	AddHeader(c)

	// 授权验证
	err := EncryptAuth(c)
	if err != nil {
		return c.JSON(result.ErrAuth.WithData(mStr.ToStr(err)))
	}

	return c.Next()
}

func AddHeader(c *fiber.Ctx) error {
	c.Set("Data-Path", "CoinMarket.net")
	return nil
}
