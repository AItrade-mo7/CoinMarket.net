package main

import (
	_ "embed"
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/ready"
	"CoinMarket.net/server/router"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)

	// 初始化系统参数
	global.Start()

	ready.Start()

	router.Start()

	// OrganizeDatabase()
}

func OrganizeDatabase() {
	fmt.Println("整理数据库")
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("BTCUSDT")
	defer db.Close()
	cur, err := db.Table.Find(db.Ctx, bson.D{})
	if err != nil {
		fmt.Println("数据库错误")
		return
	}

	for cur.Next(db.Ctx) {
		var result mOKX.TypeKd
		cur.Decode(&result)

		fmt.Println(result.TimeStr)
	}
}
