package kdata

import (
	"io/ioutil"

	"CoinMarket.net/server/analyse"
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/restApi"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

var KdataList []okxInfo.Kd

func GetKdata(InstID string) []okxInfo.Kd {
	SWAP_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	KdataList = []okxInfo.Kd{}
	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path:   "/api/v5/market/candles",
		Method: "get",
		Data: map[string]any{
			"instId": InstID,
			"bar":    "15m",
			"after":  mTime.GetUnix(),
			"limit":  300,
		},
	})

	// 本地模式
	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(SWAP_file)
	}

	if err != nil {
		global.KdataLog.Println(InstID, err)
		return nil
	}
	var result okxInfo.ReqType
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.KdataLog.Println(InstID, "Err", result)
		return nil
	}

	FormatKdata(result.Data, InstID)

	// 写入数据文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
	return KdataList
}

func FormatKdata(data any, InstID string) {
	var list []okxInfo.CandleDataType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)
	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]

		kdata := okxInfo.Kd{
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

func Storage(kdata okxInfo.Kd) {
	new_Kdata := analyse.NewKdata(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)
	go global.KdataLog.Println(mJson.JsonFormat(mJson.ToJson(new_Kdata)))
}
