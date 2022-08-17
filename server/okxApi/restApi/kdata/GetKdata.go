package kdata

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

var KdataList []mOKX.TypeKd

func GetKdata(InstID string) []mOKX.TypeKd {
	InstInfo := GetInstInfo(InstID)
	KdataList = []mOKX.TypeKd{}

	if InstInfo.InstID != InstID {
		return KdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

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
		global.LogErr("kdata.GetKdata", InstID, err)
		return nil
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("kdata.GetKdata", InstID, "Err", result)
		return nil
	}

	FormatKdata(result.Data, InstInfo)

	if len(KdataList) != 300 {
		global.KdataLog.Println("kdata.GetKdata resData", mStr.ToStr(resData))
	}

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
}

func FormatKdata(data any, Inst mOKX.TypeInst) {
	var list []mOKX.CandleDataType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	global.KdataLog.Println("kdata.FormatKdata", len(list), Inst.InstID)

	CcyName := Inst.InstID
	if Inst.InstType == "SWAP" {
		CcyName = strings.Replace(Inst.InstID, config.SWAP_suffix, "", -1)
	}
	if Inst.InstType == "SPOT" {
		CcyName = strings.Replace(Inst.InstID, config.SPOT_suffix, "", -1)
	}

	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]

		kdata := mOKX.TypeKd{
			InstID:   Inst.InstID,
			CcyName:  CcyName,
			TickSz:   Inst.TickSz,
			InstType: Inst.InstType,
			CtVal:    Inst.CtVal,
			MinSz:    Inst.MinSz,
			MaxMktSz: Inst.MaxMktSz,
			Time:     mTime.MsToTime(item[0], "0"),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
			VolCcy:   item[6],
			DataType: "GetKdata",
		}
		StorageKdata(kdata)
	}
}

func StorageKdata(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.NewKD(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)

	// global.KdataLog.Println(mJson.JsonFormat(mJson.ToJson(new_Kdata)))
}

func GetInstInfo(InstID string) (resData mOKX.TypeInst) {
	resData = mOKX.TypeInst{}

	for _, item := range okxInfo.SPOT_inst {
		if item.InstID == InstID {
			resData = item
		}
	}

	for _, item := range okxInfo.SWAP_inst {
		if item.InstID == InstID {
			resData = item
		}
	}

	return resData
}
