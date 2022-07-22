package dbUser

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/global/dbData"
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mMongo"
	"go.mongodb.org/mongo-driver/bson"
)

type NewUserOpt struct {
	Email  string
	UserID string
}

type AccountType struct {
	UserID      string `bson:"UserID"`
	AccountData dbType.AccountTable
	DB          *mMongo.DB
}

func NewUserDB(opt NewUserOpt) (resData *AccountType, resErr error) {
	resData = &AccountType{}
	resErr = nil

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("Account")

	resData.DB = db

	err := db.Ping()
	if err != nil {
		db.Close()
		errStr := fmt.Errorf("用户数据读取失败,数据库连接错误 %+v", err)
		global.LogErr(errStr)
		resErr = errStr
		return
	}

	FK := bson.D{{
		Key:   "Email",
		Value: opt.Email,
	}}
	if len(opt.UserID) > 3 {
		FK = bson.D{{
			Key:   "UserID",
			Value: opt.UserID,
		}}
	}

	var result dbType.AccountTable
	db.Table.FindOne(db.Ctx, FK).Decode(&result)

	resData.UserID = result.UserID
	resData.AccountData = result

	if len(result.UserID) > 6 {
		dbData.UserInfo = result
	}

	return
}
