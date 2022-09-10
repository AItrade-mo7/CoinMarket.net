package okxInfo

import (
	"CoinMarket.net/server/utils/dbSearch"
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

// 综合榜单数据
var TickerList []mOKX.TypeTicker

// 按照涨跌百分比排序的综合榜单数据
var TickerU_R24 []mOKX.TypeTicker

// K线数据 榜单币种 近 300 条 15 分钟间隔 共 75 小时
var MarketKdata map[string][]mOKX.TypeKd

// 单个币种的分析结果
var TickerAnalySingle map[string][]mOKX.AnalySliceType

// 榜单数据分析
var TickerAnalyWhole []mOKX.TypeWholeTickerAnaly

// 最终的分析结果 0 无结果，1 上涨 -1 下跌 ， 2  震荡上涨， -2 震荡下跌
var WholeDir = 0

// 首页的 300 条数据
var AnalyList_Client dbSearch.PagingType
var AnalyList_Serve dbSearch.PagingType
