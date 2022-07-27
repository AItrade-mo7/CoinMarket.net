package kdata

import (
	"io/ioutil"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type CandleDataType [7]string

func GetKdata(InstID string) {
	SWAP_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

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
		return
	}
	var result okxInfo.ReqType
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.KdataLog.Println(InstID, "Err", result)
		return
	}

	FormatKdata(result.Data)

	// 写入数据文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}

func FormatKdata(data any) {
	var list []CandleDataType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	// for _, item := range list {
	// }
}
