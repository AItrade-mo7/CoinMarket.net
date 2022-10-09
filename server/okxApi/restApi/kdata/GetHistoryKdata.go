package kdata

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type History300Param struct {
	InstID  string `bson:"InstID"`
	Current int    `bson:"Current"` // 以当前时间为基点往前几页
	After   int64  `bson:"After"`   // 时间 默认为当前时间
}

func GetHistory300List(opt History300Param) []mOKX.TypeKd {
	InstInfo := GetInstInfo(opt.InstID)
	History300List := []mOKX.TypeKd{}
	if InstInfo.InstID != opt.InstID {
		return History300List
	}

	now := mTime.GetUnix()
	if opt.After > 0 {
		now = mStr.ToStr(opt.After)
	}
	m300 := mCount.Mul(mStr.ToStr(mTime.UnixTimeInt64.Minute*15), "300")
	mAfter := mCount.Mul(m300, mStr.ToStr(opt.Current))
	after := mCount.Sub(now, mAfter)
	timeObj := mTime.MsToTime(after, "0")

	kdata_list := []mOKX.TypeKd{}

	kdata_3 := GetHistoryKdata(HistoryKdataParam{
		InstID:  opt.InstID,
		Current: 2,
		After:   mTime.ToUnixMsec(timeObj),
		Size:    100,
	})
	kdata_list = append(kdata_list, kdata_3...)

	kdata_2 := GetHistoryKdata(HistoryKdataParam{
		InstID:  opt.InstID,
		Current: 1,
		After:   mTime.ToUnixMsec(timeObj),
		Size:    100,
	})
	kdata_list = append(kdata_list, kdata_2...)

	kdata_1 := GetHistoryKdata(HistoryKdataParam{
		InstID:  opt.InstID,
		Current: 0,
		After:   mTime.ToUnixMsec(timeObj),
		Size:    100,
	})
	kdata_list = append(kdata_list, kdata_1...)

	CheckTicker(kdata_list)

	return History300List
}

func CheckTicker(KdataList []mOKX.TypeKd) {
	for key, val := range KdataList {
		pre := key - 1
		if pre < 0 {
			pre = 0
		}
		fmt.Println(key, val.InstID, val.TimeStr, val.TimeUnix-KdataList[pre].TimeUnix)
	}
	fmt.Println("结束")
}

type HistoryKdataParam struct {
	InstID  string `bson:"InstID"`
	Current int    `bson:"Current"` // 当前页码 0 为
	After   int64  `bson:"After"`   // 时间 默认为当前时间
	Size    int    `bson:"Size"`    // 数量 默认为100
}

func GetHistoryKdata(opt HistoryKdataParam) []mOKX.TypeKd {
	InstInfo := GetInstInfo(opt.InstID)
	HistoryKdataKdataList := []mOKX.TypeKd{}

	if InstInfo.InstID != opt.InstID {
		return HistoryKdataKdataList
	}
	Kdata_file := mStr.Join(config.Dir.JsonData, "/", opt.InstID, "-", opt.Current, "_History.json")

	now := mTime.GetUnix()
	if opt.After > 0 {
		now = mStr.ToStr(opt.After)
	}
	m100 := mCount.Mul(mStr.ToStr(mTime.UnixTimeInt64.Minute*15), mStr.ToStr(opt.Size))
	mAfter := mCount.Mul(m100, mStr.ToStr(opt.Current))
	after := mCount.Sub(now, mAfter)

	size := 100
	if opt.Size > 0 {
		size = opt.Size
	}

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/history-candles",
		Data: map[string]any{
			"instId": opt.InstID,
			"bar":    "15m",
			"after":  after,
			"limit":  size,
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
		DataType: "HistoryKdata",
	})

	global.KdataLog.Println("kdata.GetHistoryKdata", len(HistoryKdataKdataList), opt.InstID)

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return HistoryKdataKdataList
}
