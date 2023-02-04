package okxApi

import (
	"strings"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/binanceApi"
	"CoinMarket.net/server/okxApi/restApi/tickers"
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

func SetTicker() {
	binanceApi.GetTicker() // 获取 okxInfo.BinanceTickerList
	tickers.GetTicker()    // 获取 okxInfo.OKXTickerList

	if len(okxInfo.Inst) < 10 {
		global.LogErr("ready.SetTicker okxInfo.Inst 数据条目不正确", len(okxInfo.Inst))
		return
	}

	if len(okxInfo.BinanceTickerList) < 6 || len(okxInfo.OKXTickerList) < 6 {
		global.LogErr("ready.SetTicker TickerList 数据条目不正确", len(okxInfo.BinanceTickerList), len(okxInfo.OKXTickerList))
		return
	}

	tickerList := []mOKX.TypeTicker{}

	for _, okx := range okxInfo.OKXTickerList {
		for _, binance := range okxInfo.BinanceTickerList {
			if okx.InstID == binance.InstID {
				ticker := TickerCount(okx, binance)
				if len(ticker.InstID) > 2 {
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

	SWAP := mOKX.TypeInst{}
	SPOT := mOKX.TypeInst{}

	if len(Ticker.InstID) > 3 {
		SWAP = okxInfo.Inst[mStr.Join(Ticker.CcyName, config.SWAP_suffix)]
		SPOT = okxInfo.Inst[Ticker.InstID]
	}

	if len(SWAP.InstID) < 3 || len(SPOT.InstID) < 3 {
		global.LogErr("ready.TickerCount 数量不足", len(SWAP.InstID), len(SPOT.InstID))
		return mOKX.TypeTicker{}
	}

	// 合约 上架小于 36 天的不计入榜单
	diffOnLine := mCount.Sub(mStr.ToStr(Ticker.TimeUnix), SWAP.ListTime)
	diffDay := mCount.Div(diffOnLine, mTime.UnixTime.Day)
	if mCount.Le(diffDay, "36") < 0 {
		return mOKX.TypeTicker{}
	}

	return Ticker
}
