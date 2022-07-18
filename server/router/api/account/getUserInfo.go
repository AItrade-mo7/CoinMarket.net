package account

import (
	"CoinMarket.net/server/global/apiType"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/middle"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func GetUserInfo(c *fiber.Ctx) error {
	userID, err := middle.TokenAuth(c)
	if err != nil {
		return c.JSON(result.ErrToken.WithData(mStr.ToStr(err)))
	}

	if userID != dbType.UserInfo.UserID {
		return c.JSON(result.ErrToken.WithData("Token验证失败"))
	}

	var userinfoData apiType.UserInfo
	jsonStr := mJson.ToJson(dbType.UserInfo)
	jsoniter.Unmarshal(jsonStr, &userinfoData)

	return c.JSON(result.Succeed.WithData(userinfoData))
}
