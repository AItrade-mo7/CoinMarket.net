package testHunter

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*

	List := mOKX.GetKdata(mOKX.GetKdataOpt{
		InstID: "BTC-USDT",
		After:  mTime.TimeParse(mTime.Lay_ss, "2023-01-01T00:00:00"),
		Page:   0,
	})

	for _, KD := range List {
		fmt.Println(KD.TimeStr, KD.C)
	}

*/

type TestOpt struct {
	StartTime int64
	EndTime   int64
	InstID    string
}

type TestObj struct {
	StartTime int64
	EndTime   int64
	InstID    string // BTC-USDT
	KdataList []mOKX.TypeKd
}

func New(opt TestOpt) *TestObj {
	obj := TestObj{}

	NowTime := mTime.GetUnixInt64()
	earliest := mTime.TimeParse(mTime.Lay_ss, "2020-02-01T23:00:00")

	obj.EndTime = opt.EndTime
	obj.StartTime = opt.StartTime
	obj.InstID = opt.InstID

	if obj.EndTime > NowTime {
		obj.EndTime = NowTime
	}

	if obj.StartTime < earliest {
		obj.StartTime = earliest
	}

	global.Run.Println("新建回测", mJson.Format(obj))

	return &obj
}

func (_this *TestObj) StuffDBKdata(CallBack func(mOKX.TypeKd)) error {
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
		CallBack(result)
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

	db.Close()
	return nil
}
