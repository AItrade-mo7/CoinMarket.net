package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/tmpl"
	"github.com/EasyGolang/goTools/mBinance"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetBinancePosition() {
	nowBinancePosition := binanceApi.GetAccount() // 存储到数据库 BinancePosition
	global.Run.Println("读取一次 币安持仓", mJson.ToStr(nowBinancePosition))

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "AITrade",
		Timeout:  100,
	}).Connect().Collection("BinancePosition")
	defer global.Run.Println("关闭数据库连接", "BinancePosition")
	defer db.Close()
	// 先读取最近的10条数据 , 第 0 条 为 最新数据
	dbPositionList := ReadBinancePosition10(db)

	// 获取数据库最新数据
	var dbLast mBinance.PositionType
	if len(dbPositionList) > 0 {
		dbLast = dbPositionList[0]
	}
	// 用当前数据比对最新数据
	isChange := false                         // 是否存在改变
	if nowBinancePosition.Dir != dbLast.Dir { // 方向不同则改变
		isChange = true
	}
	if nowBinancePosition.InstID != dbLast.InstID { // 币种不同则改变
		isChange = true
	}

	if isChange { // 如果改变则插入最新一条并发送通知
		InsertNowBinancePosition(db, nowBinancePosition)
		NotificationChange(nowBinancePosition)
	} else { // 如果未改变，则更新当前
		UpdateBinancePosition(db, nowBinancePosition, dbLast)
	}

	// 最后，重新读取数据 并 验证
	nowPositionList := ReadBinancePosition10(db)
	// 如果数据小于 1 则报错
	if len(nowPositionList) < 1 {
		db.Close()
		resErr := fmt.Errorf("nowPositionList 读取失败 %+v", len(nowPositionList))
		global.LogErr(resErr)
	}
	dbNow := nowPositionList[0]
	isAlike := true                          // 是否相同
	if nowBinancePosition.Dir != dbNow.Dir { // 方向不同则 false
		isAlike = false
	}
	if nowBinancePosition.InstID != dbNow.InstID { // 币种不同则报错
		isAlike = false
	}
	if nowBinancePosition.CreateTime != dbNow.CreateTime { // 创建时间不同则报错
		isAlike = false
	}

	if isAlike { // 对比一致则插入内存
		okxInfo.BinancePositionList = []mBinance.PositionType{}
		okxInfo.BinancePositionList = nowPositionList
	} else { // 对比不一致则报错
		db.Close()
		resErr := fmt.Errorf("两则数据对比失败%+v %+v", mJson.ToStr(dbNow), mJson.ToStr(nowBinancePosition))
		global.LogErr(resErr)
	}
}

// 插入一条
func InsertNowBinancePosition(db *mMongo.DB, BinancePosition mBinance.PositionType) {
	_, err := db.Table.InsertOne(db.Ctx, BinancePosition)
	global.Run.Println("插入币安持仓")
	if err != nil {
		db.Close()
		resErr := fmt.Errorf("币安数据,插入数据失败 %+v", err)
		global.LogErr(resErr)
		return
	}
}

// 更新当前
func UpdateBinancePosition(db *mMongo.DB, nowData, dbLast mBinance.PositionType) {
	FK := bson.D{{
		Key:   "CreateTime",
		Value: dbLast.CreateTime,
	}}
	nowData.UpdateTime = mTime.GetUnixInt64()
	nowData.UpdateTimeStr = mTime.UnixFormat(nowData.UpdateTime)
	UK := bson.D{}
	mStruct.Traverse(nowData, func(key string, val any) {
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
	_, err := db.Table.UpdateOne(db.Ctx, FK, UK)
	global.Run.Println("更新币安持仓")
	if err != nil {
		db.Close()
		resErr := fmt.Errorf("币安数据,更新数据失败 %+v", err)
		global.LogErr(resErr)
		return
	}
}

// 读取最近10条
func ReadBinancePosition10(db *mMongo.DB) (PositList []mBinance.PositionType) {
	findOpt := options.Find()
	findOpt.SetAllowDiskUse(true)
	findOpt.SetLimit(10)
	findOpt.SetSort(map[string]int{
		"CreateTime": -1,
	})
	findFK := bson.D{}
	cursor, err := db.Table.Find(db.Ctx, findFK, findOpt)

	var nowList []mBinance.PositionType
	for cursor.Next(db.Ctx) {
		var Posit mBinance.PositionType
		cursor.Decode(&Posit)
		nowList = append(nowList, Posit)
	}

	PositList = nowList

	if err != nil {
		db.Close()
		resErr := fmt.Errorf("数据读取失败, %+v", err)
		global.LogErr(resErr)
		return
	}
	return
}

func NotificationChange(nowData mBinance.PositionType) {
	inst := okxInfo.Inst[nowData.InstID]

	dir := `<span style="color: #A69B9B;"> 错误 <span>`

	if nowData.Dir > 0 {
		dir = `<span style="color: #116F19;"> 上涨 <span>`
	}
	if nowData.Dir < 0 {
		dir = `<span style="color: #D93424;"> 下跌 <span>`
	}

	msg := mStr.Join(
		"侦测币种: ", inst.Uly, "<br />",
		"趋势方向: ", dir, "<br />",
		"数据时间: ", nowData.CreateTimeStr, "<br />",
	)

	Email := global.Email(global.EmailOpt{
		To:       config.Email.To,
		Subject:  "市场方向变更",
		Template: tmpl.SysEmail,
		SendData: tmpl.SysParam{
			Message: mJson.JsonFormat(mJson.ToJson(msg)),
			SysTime: mTime.IsoTime(),
		},
	})
	Email.Send()
}
