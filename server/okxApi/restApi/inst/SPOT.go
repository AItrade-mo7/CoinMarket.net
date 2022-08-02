package inst

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

// 获取可交易现货列表
func SPOT() {
	SPOT_file := mStr.Join(config.Dir.JsonData, "/SPOT.json")

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/public/instruments",
		Data: map[string]any{
			"instType": "SPOT",
		},
		Method:        "get",
		LocalJsonPath: SPOT_file,
		IsLocalJson:   config.AppEnv.RunMod == 1,
	})
	if err != nil {
		global.InstLog.Println("SPOT", err)
		return
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.InstLog.Println("SPOT-err", result)
		return
	}

	setSPOT_list(result.Data)

	// 写入数据文件
	go mFile.Write(SPOT_file, mStr.ToStr(resData))
}

func setSPOT_list(data any) {
	var list []mOKX.TypeInst
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	for _, val := range list {
		find := strings.Contains(val.InstID, config.SPOT_suffix)
		if find && val.State == "live" {
			SPOT_list[val.InstID] = val
		}
	}
}
