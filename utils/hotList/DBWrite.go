package hotList

import (
	"CoinMarket.net/global"
	"CoinMarket.net/okxInfo"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type ListSum struct {
	Amount24Hot []okxInfo.HotInfo
	U_R24Hot    []okxInfo.HotInfo
	U_R24AbsHot []okxInfo.HotInfo
}

type DBWriteBody struct {
	Amount24Hot []okxInfo.HotInfo `bson:"Amount24Hot"`
	U_R24Hot    []okxInfo.HotInfo `bson:"U_R24Hot"`
	U_R24AbsHot []okxInfo.HotInfo `bson:"U_R24AbsHot"`
	UpdateTime  int64             `bson:"UpdateTime"`
	RunMod      int               `bson:"RunMod"`
	Desc        string            `bson:"Desc"`
}

func DBWrite(sum ListSum) {
	// 在这里存入数据库

	if len(sum.Amount24Hot) < 2 {
		global.LogErr("hotList.DBWrite, 存储数据不足！", mStr.ToStr(mJson.ToJson(sum)))
		return
	}

	db := mMongo.New(mMongo.Opt{
		UserName: global.ServerEnv.MongoUserName,
		Password: global.ServerEnv.MongoPassword,
		Address:  global.ServerEnv.MongoAddress,
		DBName:   "Hunter",
	}).Connect().Collection("HotList")

	err := db.Ping()
	if err != nil {
		global.LogErr("hotList.DBWrite, 数据库连接错误", err)
		db.Close()
		return
	}

	var IK DBWriteBody
	IK.Amount24Hot = sum.Amount24Hot
	IK.U_R24Hot = sum.U_R24Hot
	IK.U_R24AbsHot = sum.U_R24AbsHot
	IK.UpdateTime = mTime.GetUnixInt64()

	IK.RunMod = global.ServerEnv.RunMod
	IK.Desc = "正常存储"

	InsertOneResult, err := db.Table.InsertOne(db.Ctx, IK)
	if err != nil {
		global.LogErr("hotList.DBWrite, 插入数据失败", err)
		db.Close()
		return
	}
	global.LogHotList.Println("hotList.DBWrite, 已插入数据库", mStr.ToStr(InsertOneResult))
	db.Close()
}
