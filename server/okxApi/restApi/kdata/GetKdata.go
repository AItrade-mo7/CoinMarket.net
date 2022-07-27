package kdata

import (
	"fmt"
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

	FormatKdata(result.Data, InstID)

	// 写入数据文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}

var KdataList []okxInfo.Kd

func FormatKdata(data any, InstID string) {
	var list []okxInfo.CandleDataType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)
	for _, item := range list {
		kdata := okxInfo.Kd{
			InstID:   InstID,
			TimeUnix: item[0],
			Time:     mTime.MsToTime(item[0], "0"),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
		}
		fmt.Println(kdata)
	}
}
