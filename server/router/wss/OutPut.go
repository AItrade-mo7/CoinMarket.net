package wss

import (
	"github.com/EasyGolang/goTools/mTime"
)

type OutPut struct {
	SysTime    int64  `bson:"SysTime"`    // 系统时间
	DataSource string `bson:"DataSource"` // 数据来源
}

func GetOutPut() (resData OutPut) {
	resData = OutPut{}

	resData.SysTime = mTime.GetUnixInt64()
	resData.DataSource = "CoinMarket.net"

	return
}
