package api

import (
	"fmt"

	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type GetKdataPageParam struct {
	InstID  string `bson:"InstID"`
	Current int64  `bson:"Current"` // 当前页码 0 为
}

// 获取当前页码的币种数据，并进行存储，15分钟为限额
func GetKdataPage(c *fiber.Ctx) error {
	var json GetKdataPageParam
	mFiber.Parser(c, &json)
	if len(json.InstID) < 3 {
		return c.JSON(result.Fail.WithData("需要 InstID"))
	}
	if json.Current < 0 {
		json.Current = 0
	}

	fmt.Println(json)

	return c.JSON(result.Succeed.WithData("接口开发中"))
}
