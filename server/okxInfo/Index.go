package okxInfo

import "github.com/EasyGolang/goTools/mOKX"

// 产品信息 - 现货
var SPOT_inst = make(map[string]mOKX.InstType)

// 产品信息 - 合约
var SWAP_inst = make(map[string]mOKX.InstType)

// Binance 的榜单数据
var BinanceTickerList []mOKX.BinanceTickerType // 币安的Ticker 排行

// OKX 的榜单数据
var OKXTickerList []mOKX.OKXTickerType // okx的Ticker

// 综合榜单数据
var TickerList []mOKX.TickerType

// 按照涨跌百分比排序的综合榜单数据
var TickerU_R24 []mOKX.TickerType

// K线数据 榜单币种 近 300 条 15 分钟间隔 共 75 小时
var MarketKdata = map[string][]mOKX.Kd{}

// 单个币种的分析结果
var TickerAnalyseSingle = map[string]mOKX.AnalyseSingleType{}

// 榜单数据分析
var TickerAnalyseWhole mOKX.WholeTickerAnalyseType
