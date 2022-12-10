package kdata

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type HistoryKdataParam struct {
	InstID  string `bson:"InstID"`
	Current int    `bson:"Current"` // 当前页码 0 为
	After   int64  `bson:"After"`   // 时间 默认为当前时间
	Bar     string `bson:"Bar"`
}

func GetHistoryKdata(opt HistoryKdataParam) []mOKX.TypeKd {
	InstInfo := okxInfo.Inst[opt.InstID]
	HistoryKdataKdataList := []mOKX.TypeKd{}

	if len(InstInfo.InstID) < 3 {
		return HistoryKdataKdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", opt.InstID, "-", opt.Current, "_History.json")

	Size := config.KdataLen

	now := mTime.GetUnix()
	if opt.After > 0 {
		now = mStr.ToStr(opt.After)
	}
	m100 := mCount.Mul(mStr.ToStr(mTime.UnixTimeInt64.Minute*15), mStr.ToStr(Size))
	mAfter := mCount.Mul(m100, mStr.ToStr(opt.Current))
	after := mCount.Sub(now, mAfter)

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/history-candles",
		Data: map[string]any{
			"instId": opt.InstID,
			"bar":    "15m",
			"after":  after,
			"limit":  Size,
		},
		Method:        "get",
		LocalJsonPath: Kdata_file,
		IsLocalJson:   false,
	})
	if err != nil {
		global.LogErr("kdata.GetHistoryKdata Err", opt.InstID, err)
		return nil
	}

	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("kdata.GetHistoryKdata Err", opt.InstID, result)
		return nil
	}

	HistoryKdataKdataList = mOKX.FormatKdata(mOKX.FormatKdataParam{
		Data:     result.Data,
		Inst:     InstInfo,
		DataType: "OKXKdata",
	})

	if len(HistoryKdataKdataList) > 3 {
		global.KdataLog.Println("kdata.GetHistoryKdata", len(HistoryKdataKdataList), InstInfo.InstID, HistoryKdataKdataList[0].TimeStr, HistoryKdataKdataList[len(HistoryKdataKdataList)-1].TimeStr)
	} else {
		global.KdataLog.Println("kdata.GetHistoryKdata Err", len(HistoryKdataKdataList), InstInfo.InstID)
	}

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return HistoryKdataKdataList
}
