package middle

import (
	"errors"
	"strings"

	"CoinMarket.net/server/global/config"
	"github.com/gofiber/fiber/v2"
)

func EncryptAuth(c *fiber.Ctx) error {
	EncStr := c.Get("Auth-Encrypt")
	shaStr := config.Encrypt(c.Path())
	isFind := strings.Contains(shaStr, EncStr)
	if !isFind {
		return errors.New("授权验证错误")
	}

	return nil
}
