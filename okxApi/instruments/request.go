package instruments

import (
	"strings"

	"CoinMarket.net/global"
	"CoinMarket.net/okxApi"
	"CoinMarket.net/okxInfo"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type ReqType struct {
	Code string                 `bson:"Code"`
	Data []okxInfo.InstInfoType `bson:"Data"`
	Msg  string                 `bson:"Msg"`
}

// SPOT   SWAP
func Start(instType string) {
	var reqData []byte

	// 本地测试
	if global.ServerEnv.RunMod == 1 {
		reqData = mFile.ReadFile(mStr.Join(
			"./jsonData/inst-",
			instType,
			".json",
		))
	}

	if global.ServerEnv.RunMod == 0 {
		reqData = mFetch.NewHttp(mFetch.HttpOpt{
			Origin: okxApi.BaseUrl.Rest,
			Path:   "/api/v5/public/instruments",
			Data: map[string]any{
				"instType": instType,
			},
		}).Get()
	}

	if len(reqData) < 30 {
		global.LogInst.Println(instType, "获取 HotList 数据 -- 失败")
		return
	} else {
		global.LogInst.Println(instType, "获取 instruments 数据")
	}

	var data ReqType
	err := jsoniter.Unmarshal(reqData, &data)
	if err != nil {
		global.LogErr(instType, "instruments 数据格式化失败 : "+mStr.ToStr(reqData))
		return
	}

	if data.Code != "0" {
		global.LogErr(instType, "instruments data.code 出现问题 : "+mStr.ToStr(reqData))
		return
	}

	if len(data.Data) < 10 {
		global.LogErr(instType, "instruments 可用数量不足 : "+mStr.ToStr(reqData))
		return
	}

	list := data.Data

	okxInfo.ClearInst(instType)
	SetInst(list)
}

func SetInst(data []okxInfo.InstInfoType) {
	liveList := []okxInfo.InstInfoType{}
	preopenList := []okxInfo.InstInfoType{}

	for _, val := range data {
		if val.InstType == "SPOT" {
			// 现货 计价货币 不为 USDT 跳过
			if val.QuoteCcy != "USDT" {
				continue
			}
		}
		if val.InstType == "SWAP" {
			// 合约 盈亏结算和保证金币种 不为 USDT 则跳过
			if val.SettleCcy != "USDT" {
				continue
			}
		}

		find := strings.Contains(val.InstID, "-USDT") // 只保留 USDT
		// InstID 如果不包含 -USDT 则跳过
		if !find {
			continue
		}

		// 交易中的
		if val.State == "live" {
			liveList = append(liveList, val)
		}

		// 预上线的
		if val.State == "preopen" {
			preopenList = append(preopenList, val)
		}
	}

	if len(preopenList) > 1 {
		first := preopenList[0]

		if first.InstType == "SWAP" {
			okxInfo.PreopenInstList.SWAP = preopenList
		}

		if first.InstType == "SPOT" {
			okxInfo.PreopenInstList.SPOT = preopenList
		}

	}

	if len(liveList) > 1 {
		first := liveList[0]

		if first.InstType == "SWAP" {
			okxInfo.InstList.SWAP = liveList
		}

		if first.InstType == "SPOT" {
			okxInfo.InstList.SPOT = liveList
		}

	} else {
		global.LogErr("instruments live 可用长度不足")
		return
	}
}
