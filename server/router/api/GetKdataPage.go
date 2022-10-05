package api

import (
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type GetKdataPageParam struct {
	InstID  string `json:"InstID"`  // 列表
	Current int64  `bson:"Current"` // 当前页码
}

// 获取当前页码的币种数据，并进行存储，15分钟为限额
func GetKdataPage(c *fiber.Ctx) error {
	var json GetKdataPageParam
	mFiber.Parser(c, &json)

	return c.JSON(result.Succeed.WithData("接口开发中"))
}
