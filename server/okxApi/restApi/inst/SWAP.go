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

// 获取可交易合约列表
func SWAP() {
	SWAP_file := mStr.Join(config.Dir.JsonData, "/SWAP.json")
	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/public/instruments",
		Data: map[string]any{
			"TypeInst": "SWAP",
		},
		Method:        "get",
		LocalJsonPath: SWAP_file,
		IsLocalJson:   config.AppEnv.RunMod == 1,
	})
	if err != nil {
		global.InstLog.Println("SWAP", err)
		return
	}

	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.InstLog.Println("SPOT-err", result)
		return
	}

	setSWAP_list(result.Data)

	// 写入日志文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}

func setSWAP_list(data any) {
	var list []mOKX.TypeInst
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	for _, val := range list {

		find := strings.Contains(val.InstID, config.SWAP_suffix) // 统一计价单位
		if find && val.State == "live" {
			SWAP_list[val.InstID] = val
		}
	}
}
