package account

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"CoinMarket.net/server/router/result"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var json struct {
		Email    string `bson:"Email"`
		Password string `bson:"Password"`
	}
	mFiber.Parser(c, &json)

	if json.Email != dbType.UserInfo.Email {
		return c.JSON(result.ErrAccount.WithMsg("该账号未注册，请先去注册"))
	}

	if json.Password != dbType.UserInfo.Password {
		return c.JSON(result.ErrPassword.WithMsg("密码错误"))
	}

	var loginSucceedData struct {
		Email    string `bson:"Email"`    // 用户 Email
		UserID   string `bson:"UserID"`   // 用户 ID
		Avatar   string `bson:"Avatar"`   // 用户头像
		NickName string `bson:"NickName"` // 用户昵称
		Token    string `bson:"Token"`
	}

	loginSucceedData.Email = dbType.UserInfo.Email
	loginSucceedData.UserID = dbType.UserInfo.UserID
	loginSucceedData.Avatar = dbType.UserInfo.Avatar
	loginSucceedData.NickName = dbType.UserInfo.NickName
	loginSucceedData.Token = mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: config.SecretKey,
		ExpiresAt: time.Now().Add(72 * time.Hour),
		Message:   dbType.UserInfo.UserID,
		Issuer:    "FileServe.net",
		Subject:   "UserToken",
	}).Generate()

	return c.JSON(result.RightLogin.WithData(loginSucceedData))
}
