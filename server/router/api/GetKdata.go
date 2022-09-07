package api

import (
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

type GetAnalyKdataParam struct {
	InstID string `json:"InstID"` // 列表
}

func GetKdata(c *fiber.Ctx) error {
	var json GetAnalyKdataParam
	mFiber.Parser(c, &json)

	if len(json.InstID) < 3 {
		return c.JSON(result.Fail.WithData("需要 InstID"))
	}

	KdataList := okxInfo.MarketKdata[json.InstID]

	if len(KdataList) < 120 {
		return c.JSON(result.Fail.WithData(mStr.Join("长度不足,", len(KdataList))))
	}

	return c.JSON(result.Succeed.WithData(KdataList))
}
