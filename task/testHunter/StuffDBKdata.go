package testHunter

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (_this *TestObj) StuffDBKdata() error {
	total := (_this.EndTime - _this.StartTime) / mTime.UnixTimeInt64.Hour
	if total < 1 {
		return fmt.Errorf("total 数量太少")
	}
	Timeout := int(total) * 10
	if Timeout < 100 {
		Timeout = 100
	}

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "CoinMarket",
		Timeout:  Timeout,
	}).Connect().Collection(_this.InstID)
	defer db.Close()
	findOpt := options.Find()
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})
	findOpt.SetAllowDiskUse(true)
	FK := bson.D{}
	FK = append(FK, bson.E{
		Key: "TimeUnix",
		Value: bson.D{
			{
				Key:   "$gte", // 大于或等于
				Value: _this.StartTime,
			}, {
				Key:   "$lte", // 小于或等于
				Value: _this.EndTime,
			},
		},
	})
	cur, err := db.Table.Find(db.Ctx, FK, findOpt)
	if err != nil {
		db.Close()
		return err
	}

	AllList := []mOKX.TypeKd{}
	for cur.Next(db.Ctx) {
		var result mOKX.TypeKd
		cur.Decode(&result)
		AllList = append(AllList, result)
	}

	// 检查数据
	for key := range AllList {
		preIndex := key - 1
		if preIndex < 0 {
			preIndex = 0
		}
		preItem := AllList[preIndex]
		nowItem := AllList[key]
		global.Run.Println(nowItem.TimeUnix - preItem.TimeUnix - 3600000)
	}

	_this.KdataList = []mOKX.TypeKd{}
	_this.KdataList = AllList

	db.Close()
	return nil
}
