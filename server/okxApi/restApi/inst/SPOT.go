package inst

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

	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(SWAP_file)
	}
	if err != nil {
		global.InstLog.Println("SPOT", err)
		return
	}

	fmt.Println(resData)

	var result okxInfo.ReqType
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.InstLog.Println("SPOT-err", result)
		return
	}

	SetInst(result.Data)

	// 写入数据文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}

func SetInst(data any) {
	var list []okxInfo.InstType

	jsonStr := mJson.ToJson(data)

	jsoniter.Unmarshal(jsonStr, &list)

	fmt.Println(list)

	// for _, val := range data {

	// 	find := strings.Contains(val.InstID, "-USDT") // 只保留 USDT
	// 	if find && val.State == "live" {
	// 		// inst := InstCount(val)
	// 		// instList = append(instList, inst)
	// 	}
	// }
	// okxInfo.InstList = instList
}
