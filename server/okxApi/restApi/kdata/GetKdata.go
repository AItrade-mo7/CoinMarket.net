package kdata

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

func GetKdata(InstID string, Size int) []mOKX.TypeKd {
	InstInfo := GetInstInfo(InstID)
	KdataList := []mOKX.TypeKd{}

	if InstInfo.InstID != InstID {
		return KdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	limit := Size
	if limit < 100 {
		limit = 100
	}

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/candles",
		Data: map[string]any{
			"instId": InstID,
			"bar":    "15m",
			"after":  mTime.GetUnix(),
			"limit":  limit,
		},
		Method:        "get",
		LocalJsonPath: Kdata_file,
		IsLocalJson:   config.SysEnv.RunMod == 1,
	})
	if err != nil {
		global.LogErr("kdata.GetKdata Err", InstID, err)
		return nil
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("kdata.GetKdata Err", InstID, result)
		return nil
	}

	KdataList = mOKX.FormatKdata(mOKX.FormatKdataParam{
		Data:     result.Data,
		Inst:     InstInfo,
		DataType: "GetKdata",
	})

	global.KdataLog.Println("kdata.GetKdata", len(KdataList), InstID)

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
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
