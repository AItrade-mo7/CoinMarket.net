package testHunter

import (
	"CoinMarket.net/server/global"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
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
