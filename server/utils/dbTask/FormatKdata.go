package dbTask

import (
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FormatKdata() {
	okxApi.SetInst() // 获取并设置交易产品信息
	go SetKdata("BTC")
	SetKdata("ETH")
}

var Page = 5

func SetKdata(CcyName string) {
	tableName := CcyName + "USDT"
	InstID := CcyName + "-USDT"

	AllList := []mOKX.TypeKd{}

	for i := 0; i < Page; i++ {
		time.Sleep(time.Second / 3)
		List := okxApi.GetKdata(okxApi.GetKdataOpt{
			InstID:  InstID,
			Current: i, // 当前页码 0 为
			After:   mTime.GetUnixInt64(),
			Size:    100,
		})

		for i := len(List) - 1; i >= 0; i-- {
			AllList = append(AllList, List[i])
		}
		global.Run.Println(List[0].TimeStr, List[len(List)-1].TimeStr)
	}

	// 链接数据库
	Timeout := len(AllList) * 10
	if Timeout < 100 {
		Timeout = 100
	}
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  Timeout,
	}).Connect().Collection(tableName)
	defer global.Run.Println("关闭数据库连接", tableName)
	defer db.Close()

	for _, Kd := range AllList {
		FK := bson.D{{
			Key:   "TimeUnix",
			Value: Kd.TimeUnix,
		}}
		UK := bson.D{}
		mStruct.Traverse(Kd, func(key string, val any) {
			UK = append(UK, bson.E{
				Key: "$set",
				Value: bson.D{
					{
						Key:   key,
						Value: val,
					},
				},
			})
		})

		upOpt := options.Update()
		upOpt.SetUpsert(true)
		_, err := db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
		if err != nil {
			global.LogErr(tableName+"数据更插失败  %+v", err)
		}
	}
}
