package hotList

import (
	"CoinMarket.net/global"
	"github.com/EasyGolang/goTools/mMongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	list := hotList.DBReadeNew(3)

	fmt.Println(mJson.ToStr(list))

*/

// 获取最新的n条数据
func DBReadeNew(n int64) []DBWriteBody {
	db := mMongo.New(mMongo.Opt{
		UserName: global.ServerEnv.MongoUserName,
		Password: global.ServerEnv.MongoPassword,
		Address:  global.ServerEnv.MongoAddress,
		DBName:   "Hunter",
	}).Connect().Collection("HotList")

	var Body []DBWriteBody

	err := db.Ping()
	if err != nil {
		global.LogErr("hotList.DBReade,  数据库连接错误", err)
		db.Close()
		return Body
	}

	JK := bson.D{{}}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{
		{
			Key:   "_id",
			Value: -1,
		},
	})
	findOptions.SetLimit(n)

	FindResult, err := db.Table.Find(db.Ctx, JK, findOptions)
	if err != nil {
		global.LogErr("hotList.DBReade, 数据读取失败", err)
		db.Close()
		return Body
	}
	err = FindResult.All(db.Ctx, &Body)

	global.LogHotList.Println("hotList.DBReade, 读取结果 err :", err)
	db.Close()
	return Body
}
