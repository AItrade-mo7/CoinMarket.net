package binanceApi

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
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
	Size    int    `bson:"Size"`    // 数量 默认为100
}

var KdataList []mOKX.TypeKd

func GetKdata(opt GetKdataParam) []mOKX.TypeKd {
	KdataList = []mOKX.TypeKd{}
	Kdata_file := mStr.Join(config.Dir.JsonData, "/B-", opt.Symbol, ".json")

	limit := opt.Size
	if limit < 100 {
		limit = 100
	}

	now := mTime.GetUnix()
	if opt.After > 0 {
		now = mStr.ToStr(opt.After)
	}
	m100 := mCount.Mul(mStr.ToStr(mTime.UnixTimeInt64.Minute*15), mStr.ToStr(opt.Size))
	mAfter := mCount.Mul(m100, mStr.ToStr(opt.Current))
	after := mCount.Sub(now, mAfter)

	size := 100
	if opt.Size > 0 {
		size = opt.Size
	}

	if opt.After == 0 {
		after = ""
	}

	resData, err := mOKX.FetchBinance(mOKX.FetchBinanceOpt{
		Path:   "/api/v3/klines",
		Method: "get",
		Data: map[string]any{
			"symbol":   opt.Symbol,
			"interval": "15m",
			"endTime":  after,
			"limit":    size,
		},
		LocalJsonPath: Kdata_file,
		IsLocalJson:   config.SysEnv.RunMod == 1,
	})
	if err != nil {
		global.LogErr("binanceApi.GetKdata Err", err)
		return nil
	}

	FormatKdata(resData, opt.Symbol)

	if len(KdataList) < limit {
		global.KdataLog.Println("binanceApi.GetKdata resData", opt.Symbol, mStr.ToStr(resData))
	}

	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
}

func FormatKdata(data []byte, Symbol string) {
	var listStr [][12]any
	jsoniter.Unmarshal(data, &listStr)

	global.BinanceKdataLog.Println("binanceApi.FormatKdata", len(listStr), Symbol)

	InstID := Symbol

	for _, item := range okxInfo.TickerVol {
		if item.Symbol == Symbol {
			InstID = item.InstID
			break
		}
	}

	for _, item := range listStr {
		TimeStr := mStr.ToStr(mJson.ToJson(item[0]))

		kdata := mOKX.TypeKd{
			InstID:   InstID,
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
		Storage(kdata)
	}
}

func Storage(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.NewKD(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)
}
