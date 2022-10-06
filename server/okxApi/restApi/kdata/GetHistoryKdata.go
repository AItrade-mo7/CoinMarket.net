package kdata

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type HistoryKdataParam struct {
	InstID  string `bson:"InstID"`
	Current int64  `bson:"Current"` // 当前页码 0 为
	Size    int    `bson:"Size"`    // 数量
}

func GetHistoryKdata(opt HistoryKdataParam) []mOKX.TypeKd {
	InstInfo := GetInstInfo(opt.InstID)
	HistoryKdataKdataList := []mOKX.TypeKd{}

	if InstInfo.InstID != opt.InstID {
		return HistoryKdataKdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", opt.InstID, "-", opt.Current, "_History.json")

	now := mTime.GetUnix()
	m100 := mCount.Mul(mStr.ToStr(mTime.UnixTimeInt64.Minute*15), mStr.ToStr(opt.Size))
	mAfter := mCount.Mul(m100, mStr.ToStr(opt.Current))
	after := mCount.Sub(now, mAfter)

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/history-candles",
		Data: map[string]any{
			"instId": opt.InstID,
			"bar":    "15m",
			"after":  mStr.ToStr(after),
			"limit":  opt.Size,
		},
		Method:        "get",
		LocalJsonPath: Kdata_file,
		IsLocalJson:   false,
	})
	if err != nil {
		global.LogErr("kdata.GetHistoryKdata", opt.InstID, err)
		return nil
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("kdata.GetHistoryKdata", opt.InstID, "Err", result)
		return nil
	}

	HistoryKdataKdataList = mOKX.FormatKdata(mOKX.FormatKdataParam{
		Data:     result.Data,
		Inst:     InstInfo,
		DataType: "HistoryKdata",
	})

	if len(HistoryKdataKdataList) < 120 {
		global.KdataLog.Println("kdata.GetHistoryKdata resData", opt.InstID, mStr.ToStr(resData))
	}

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return HistoryKdataKdataList
}
