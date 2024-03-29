package middle

import (
	"errors"
	"fmt"
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func TokenAuth(c *fiber.Ctx) (UserID string, err error) {
	Message := ""
	UserID = ""
	err = nil

	Token := c.Get("Token")
	if len(Token) < 1 {
		err = errors.New("缺少Token")
		return
	}

	Claims, AuthOk := mEncrypt.ParseToken(Token, config.SecretKey)
	if !AuthOk {
		err = errors.New("Token验证失败")
		return
	}

	Message = Claims.Message
	UserID = Message
	if len(UserID) != 32 {
		err = errors.New("Token解析失败")
		return
	}

	ExpiresAt := Claims.StandardClaims.ExpiresAt
	now := time.Now().Unix()

	if ExpiresAt-now < 0 {
		err = errors.New("Token过期")
		return
	}

	// 数据库验证
	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect()
	if err != nil {
		err = fmt.Errorf("数据库连接错误")
		return
	}
	defer db.Close()
	db.Collection("VerifyToken")

	var dbRes dbType.TokenTable
	FK := bson.D{{
		Key:   "UserID",
		Value: UserID,
	}}
	db.Table.FindOne(db.Ctx, FK).Decode(&dbRes)
	if dbRes.Token != Token {
		db.Close()
		err = fmt.Errorf("Token验证失败")
		return
	}
	db.Close()

	nowUnix := mTime.GetUnixInt64()
	if nowUnix-dbRes.CreateTime > mTime.UnixTimeInt64.Day {
		err = errors.New("Token过期")
	}
	return
}
