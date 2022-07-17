package tickersAll

import (
	"fmt"

	"CoinMarket.net/global"
	"CoinMarket.net/okxApi"
	"CoinMarket.net/okxInfo"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type ReqType struct {
	Code string            `bson:"Code"`
	Msg  string            `bson:"Msg"`
	Data []okxInfo.HotInfo `bson:"Data"`
}

func Start() {
	if len(okxInfo.InstList.SWAP) < 10 && len(okxInfo.InstList.SPOT) < 10 {
		errorsStr := fmt.Errorf("产品数据不足 10 条！！")
		global.LogErr(errorsStr)
		return
	}

	var reqData []byte

	// 本地测试
	if global.ServerEnv.RunMod == 1 {
		reqData = mFile.ReadFile("./jsonData/ticker-all.json")
	}

	if global.ServerEnv.RunMod == 0 {
		reqData = mFetch.NewHttp(mFetch.HttpOpt{
			Origin: okxApi.BaseUrl.Rest,
			Path:   "/api/v5/market/tickers",
			Data: map[string]any{
				"instType": "SWAP",
			},
		}).Get()
	}

	if len(reqData) < 30 {
		global.LogHotList.Println("获取 HotList 数据 -- 失败")
		return
	} else {
		global.LogHotList.Println("获取 HotList 数据")
	}

	var data ReqType
	err := jsoniter.Unmarshal(reqData, &data)
	if err != nil {
		global.LogErr("HotList 数据格式化失败 : " + mStr.ToStr(reqData))
		return
	}

	if data.Code != "0" {
		global.LogErr("HotList data.code 出现问题 : " + mStr.ToStr(reqData))
		return
	}

	if len(data.Data) < 5 {
		global.LogErr("HotList 数据不足5条！ : " + mStr.ToStr(reqData))
		return
	}

	list := data.Data
	SetHotList(list)
}
