package candles

import (
	"fmt"
	"io/ioutil"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

func GetCandles(InstID string) {
	SWAP_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path:   "/api/v5/market/candles",
		Method: "get",
		Data: map[string]any{
			"instId": InstID,
			"bar":    "10m",
			"after":  mTime.GetUnix(),
			"limit":  300,
		},
	})
	// 本地模式
	if config.AppEnv.RunMod == 1 {
		resData, err = ioutil.ReadFile(SWAP_file)
	}

	if err != nil {
		global.InstLog.Println(InstID, err)
		return
	}
	var result okxInfo.ReqType
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.InstLog.Println(InstID, "Err", result)
		return
	}

	fmt.Println("111", result.Code)

	// 写入数据文件
	go mFile.Write(SWAP_file, mStr.ToStr(resData))
}
