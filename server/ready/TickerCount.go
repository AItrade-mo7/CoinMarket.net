package ready

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"CoinMarket.net/server/okxInfo/analyse/sort"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

func SetTicker() {
	if len(okxInfo.BinanceTickerList) != 10 || len(okxInfo.OKXTickerList) != 10 {
		global.InstLog.Println("TickerList 数据条目不正确", len(okxInfo.BinanceTickerList), len(okxInfo.OKXTickerList))
	}

	tickerList := []mOKX.TickerType{}

	for _, okx := range okxInfo.OKXTickerList {
		for _, binance := range okxInfo.BinanceTickerList {
			if okx.InstID == binance.InstID {
				ticker := TickerCount(okx, binance)
				tickerList = append(tickerList, ticker)
				break
			}
		}
	}

	VolumeSortList := sort.Volume(tickerList)
	okxInfo.TickerList = VolumeSortList
	okxInfo.TickerU_R24 = sort.U_R24(VolumeSortList)
}

func TickerCount(OKXTicker mOKX.OKXTickerType, BinanceTicker mOKX.BinanceTickerType) (Ticker mOKX.TickerType) {
	Ticker = mOKX.TickerType{}
	Ticker.InstID = OKXTicker.InstID
	Ticker.CcyName = strings.Replace(Ticker.InstID, config.SPOT_suffix, "", -1)
	Ticker.Last = OKXTicker.Last
	Ticker.Open24H = OKXTicker.Open24H
	Ticker.High24H = OKXTicker.High24H
	Ticker.Low24H = OKXTicker.Low24H
	Ticker.OKXVol24H = OKXTicker.VolCcy24H
	Ticker.BinanceVol24H = BinanceTicker.QuoteVolume
	Ticker.U_R24 = mCount.RoseCent(OKXTicker.Last, OKXTicker.Open24H)
	Ticker.Volume = mCount.Add(OKXTicker.VolCcy24H, BinanceTicker.QuoteVolume)
	Ticker.OkxVolRose = mCount.PerCent(Ticker.OKXVol24H, Ticker.Volume)
	Ticker.BinanceVolRose = mCount.PerCent(Ticker.BinanceVol24H, Ticker.Volume)
	Ticker.Ts = mTime.ToUnixMsec(mTime.MsToTime(OKXTicker.Ts, "0"))

	return Ticker
}
