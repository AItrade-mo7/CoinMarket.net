package kdata

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

var KdataList []mOKX.TypeKd

func GetKdata(InstID string) []mOKX.TypeKd {
	Kdata_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	KdataList = []mOKX.TypeKd{}
	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/candles",
		Data: map[string]any{
			"instId": InstID,
			"bar":    "15m",
			"after":  mTime.GetUnix(),
			"limit":  300,
		},
		Method:        "get",
		LocalJsonPath: Kdata_file,
		IsLocalJson:   config.AppEnv.RunMod == 1,
	})
	if err != nil {
		global.KdataLog.Println(InstID, err)
		return nil
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.KdataLog.Println(InstID, "Err", result)
		return nil
	}

	FormatKdata(result.Data, InstID)

	// 写入数据文件
	go mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
}

func FormatKdata(data any, InstID string) {
	var list []mOKX.CandleDataType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)
	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]

		kdata := mOKX.TypeKd{
			InstID:   InstID,
			Time:     mTime.MsToTime(item[0], "0"),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
			VolCcy:   item[6],
			Type:     "GetKdata",
		}
		Storage(kdata)
	}
}

func Storage(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.AnalyNewKd(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)
	global.KdataLog.Println(mJson.JsonFormat(mJson.ToJson(new_Kdata)))
}
