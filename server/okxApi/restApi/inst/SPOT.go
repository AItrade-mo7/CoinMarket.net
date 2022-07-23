package inst

import (
	"io/ioutil"
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

// 获取可交易现货列表
func SPOT() {
	SWAP_file := mStr.Join(config.Dir.JsonData, "/SPOT.json")

	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path:   "/api/v5/public/instruments",
		Method: "get",
		Data: map[string]any{
			"instType": "SPOT",
		},
	})
	// 本地模式
	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(SWAP_file)
	}

	if err != nil {
		global.InstLog.Println("SPOT", err)
		return
	}
	var result okxInfo.ReqType
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.InstLog.Println("SPOT-err", result)
		return
	}

	setSPOT_list(result.Data)

	// 写入数据文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}

func setSPOT_list(data any) {
	var list []okxInfo.InstType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	for _, val := range list {
		find := strings.Contains(val.InstID, config.SPOT_suffix)
		if find && val.State == "live" {
			SPOT_list[val.InstID] = val
		}
	}
}
