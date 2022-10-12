package okxInfo

import (
	"CoinMarket.net/server/tickerAnaly"
	"github.com/EasyGolang/goTools/mOKX"
)

// 产品信息 - 现货
var SPOT_inst map[string]mOKX.TypeInst

// 产品信息 - 合约
var SWAP_inst map[string]mOKX.TypeInst

// Binance 的榜单数据
var BinanceTickerList []mOKX.TypeBinanceTicker // 币安的Ticker 排行

// OKX 的榜单数据
var OKXTickerList []mOKX.TypeOKXTicker // okx的Ticker

// 当前综合榜单数据
var TickerVol []mOKX.TypeTicker

// 当前榜单 K线数据 榜单币种 近 300 条 15 分钟间隔 共 75 小时
var TickerKdata map[string][]mOKX.TypeKd

// 当前的分析结果
var AnalyDetail tickerAnaly.AnalyResult
