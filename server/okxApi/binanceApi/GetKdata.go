package binanceApi

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/utils"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type GetKdataParam struct {
	Symbol  string `bson:"Symbol"`
	Current int    `bson:"Current"` // 当前页码 0 为
	After   int64  `bson:"After"`   // 时间 默认为当前时间
	Bar     string `bson:"Bar"`
}

func GetKdata(opt GetKdataParam) (KdataList []mOKX.TypeKd) {
	InstInfo := okxInfo.Inst[opt.Symbol]
	KdataList = []mOKX.TypeKd{}

	if len(InstInfo.Symbol) < 3 {
		return KdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/B-", opt.Symbol, ".json")

	Size := config.KdataLen

	after := ""

	if opt.After > 0 {
		now := mStr.ToStr(opt.After)
		m100 := mCount.Mul(mStr.ToStr(mTime.UnixTimeInt64.Minute*15), mStr.ToStr(Size))
		mAfter := mCount.Mul(m100, mStr.ToStr(opt.Current))
		after = mCount.Sub(now, mAfter)
	}

	resData, err := utils.FetchBinance(utils.FetchBinanceOpt{
		Path:   "/api/v3/klines",
		Method: "get",
		Data: map[string]any{
			"symbol":   opt.Symbol,
			"interval": opt.Bar,
			"endTime":  after,
			"limit":    Size,
		},
		LocalJsonPath: Kdata_file,
		IsLocalJson:   config.SysEnv.RunMod == 1,
	})
	if err != nil {
		global.LogErr("binanceApi.GetKdata Err", err)
		return KdataList
	}

	rList := FormatKdata(resData, opt.Symbol)

	if len(rList) < Size {
		global.KdataLog.Println("binanceApi.GetKdata resData", opt.Symbol, mStr.ToStr(resData))
	}

	KdataList = rList

	if len(KdataList) > 3 {
		global.KdataLog.Println("binanceApi.GetKdata", len(KdataList), InstInfo.Symbol, KdataList[0].TimeStr, KdataList[len(KdataList)-1].TimeStr)
	} else {
		global.KdataLog.Println("binanceApi.GetKdata Err", len(KdataList), InstInfo.Symbol)
	}

	mFile.Write(Kdata_file, mStr.ToStr(resData))

	return KdataList
}

func FormatKdata(data []byte, Symbol string) (rList []mOKX.TypeKd) {
	rList = []mOKX.TypeKd{}
	var listStr [][12]any
	jsoniter.Unmarshal(data, &listStr)

	global.BinanceKdataLog.Println("binanceApi.FormatKdata", len(listStr), Symbol)

	SPOT := okxInfo.Inst[Symbol]
	if len(SPOT.InstID) < 3 {
		return
	}

	for _, item := range listStr {
		TimeStr := mStr.ToStr(mJson.ToJson(item[0]))

		kdata := mOKX.TypeKd{
			InstID:   SPOT.InstID,
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(TimeStr, "0")),
			TimeStr:  mTime.UnixFormat(TimeStr),
			O:        mStr.ToStr(item[1]),
			H:        mStr.ToStr(item[2]),
			L:        mStr.ToStr(item[3]),
			C:        mStr.ToStr(item[4]),
			Vol:      mStr.ToStr(item[5]),
			VolCcy:   mStr.ToStr(item[7]),
			DataType: "BinanceKdata",
		}
		new_Kdata := mOKX.NewKD(kdata, rList)
		rList = append(rList, new_Kdata)
	}

	return rList
}
