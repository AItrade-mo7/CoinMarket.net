package dbTidy

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mMongo"
)

func CheckCoinKdata() {
	fmt.Println("开始检查重复数据")
	CoinName := "BTC"
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  99999,
	}).Connect().Collection(CoinName + "USDT")
	defer global.Run.Println("关闭数据库连接" + CoinName)
	defer db.Close()
}
