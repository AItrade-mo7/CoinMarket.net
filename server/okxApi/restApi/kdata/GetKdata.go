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

func GetKdata(InstID string, Bar string) []mOKX.TypeKd {
	InstInfo := okxInfo.Inst[InstID]
	KdataList := []mOKX.TypeKd{}

	if len(InstInfo.InstID) < 3 {
		return KdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	limit := config.KdataLen

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/candles",
		Data: map[string]any{
			"instId": InstID,
			"bar":    Bar,
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
		DataType: "OKXKdata",
	})

	if len(KdataList) > 3 {
		global.KdataLog.Println("kdata.GetKdata", len(KdataList), InstID, KdataList[0].TimeStr, KdataList[len(KdataList)-1].TimeStr)
	} else {
		global.KdataLog.Println("kdata.GetKdata Err", len(KdataList), InstID)
	}

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
}
