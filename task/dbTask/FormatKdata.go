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

type FormatKdataObj struct {
	Page int
}

func FormatKdata() *FormatKdataObj {
	okxApi.SetInst() // 获取并设置交易产品信息
	obj := FormatKdataObj{}

	obj.Page = 270 // 请求近270 页的数据，一页 100 条。 每条 1 小时

	// 开始塞入 BTC 的数据
	obj.SetKdata("BTC")

	// 开始塞入 ETH 的数据
	obj.SetKdata("ETH")

	return &obj
}

func (obj *FormatKdataObj) SetKdata(CcyName string) {
	InstID := CcyName + "-USDT"

	AllList := []mOKX.TypeKd{}

	// 请求数据
	for i := 0; i < obj.Page; i++ {
		time.Sleep(time.Second / 3)
		List := mOKX.GetKdata(mOKX.GetKdataOpt{
			InstID: InstID,
			Page:   i,
			After:  mTime.GetUnixInt64(),
		})

		for i := len(List) - 1; i >= 0; i-- {
			AllList = append(AllList, List[i])
		}
		if len(List) > 0 {
			global.Run.Println(InstID, List[0].TimeStr, List[len(List)-1].TimeStr)
		} else {
			global.Run.Println("请求数据出错", len(List), InstID, i)
		}
	}

	// 数据检查
	for key := range AllList {
		preIndex := key - 1
		if preIndex < 0 {
			preIndex = 0
		}
		preItem := AllList[preIndex]
		nowItem := AllList[key]
		global.Run.Println(nowItem.TimeUnix - preItem.TimeUnix)
	}

	// 链接数据库
	Timeout := len(AllList) * 10
	if Timeout < 100 {
		Timeout = 100
	}
	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "CoinMarket",
		Timeout:  Timeout,
	}).Connect()
	if err != nil {
		panic(err)
	}
	defer global.Run.Println("关闭数据库连接", InstID)
	defer db.Close()
	db.Collection(InstID)

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
			global.Run.Println(InstID+"数据更插失败  %+v", err)
		}
		global.Run.Println("数据更插成功", InstID, Kd.TimeStr)
	}
}
