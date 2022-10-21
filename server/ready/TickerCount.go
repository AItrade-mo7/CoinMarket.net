package ready

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

func SetTicker() {
	if len(okxInfo.BinanceTickerList) != 15 || len(okxInfo.OKXTickerList) != 15 {
		global.LogErr("ready.SetTicker TickerList 数据条目不正确", len(okxInfo.BinanceTickerList), len(okxInfo.OKXTickerList))
	}

	tickerList := []mOKX.TypeTicker{}

	for _, okx := range okxInfo.OKXTickerList {
		for _, binance := range okxInfo.BinanceTickerList {
			if okx.InstID == binance.InstID {
				ticker := TickerCount(okx, binance)
				if len(ticker.InstID) > 4 {
					tickerList = append(tickerList, ticker)
				}
				break
			}
		}
	}

	VolumeSortList := mOKX.SortVolume(tickerList)
	okxInfo.TickerVol = []mOKX.TypeTicker{}
	okxInfo.TickerVol = VolumeSortList
}

func TickerCount(OKXTicker mOKX.TypeOKXTicker, BinanceTicker mOKX.TypeBinanceTicker) (Ticker mOKX.TypeTicker) {
	Ticker = mOKX.TypeTicker{}
	Ticker.InstID = OKXTicker.InstID
	Ticker.Symbol = BinanceTicker.Symbol
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
	Ticker.TimeUnix = mTime.ToUnixMsec(mTime.MsToTime(OKXTicker.Ts, "0"))
	Ticker.TimeStr = mTime.UnixFormat(Ticker.TimeUnix)
	Ticker.SWAP = mOKX.TypeInst{}
	Ticker.SPOT = mOKX.TypeInst{}
	if len(Ticker.InstID) > 3 {
		for _, SWAP := range okxInfo.SWAP_inst {
			if SWAP.Uly == Ticker.InstID {
				Ticker.SWAP = SWAP
				break
			}
		}
		for _, SPOT := range okxInfo.SPOT_inst {
			if SPOT.InstID == Ticker.InstID {
				Ticker.SPOT = SPOT
				break
			}
		}
	}

	// 上架小于 32 天的不计入榜单
	diffOnLine := mCount.Sub(mStr.ToStr(Ticker.TimeUnix), Ticker.SWAP.ListTime)
	if mCount.Le(diffOnLine, "32") < 0 {
		return mOKX.TypeTicker{}
	}

	return Ticker
}
