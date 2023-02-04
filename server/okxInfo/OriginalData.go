package okxInfo

import (
	"CoinMarket.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mOKX/binance"
)

// 产品信息
var Inst map[string]mOKX.TypeInst

// Binance 的榜单数据
var BinanceTickerList []mOKX.TypeBinanceTicker // 币安的Ticker 排行

// OKX 的榜单数据
var OKXTickerList []mOKX.TypeOKXTicker // okx的Ticker

// 当前综合榜单数据
var TickerVol []mOKX.TypeTicker

// 当前榜单 K线数据 榜单币种 近 300 条 15 分钟间隔 共 75 小时
var TickerKdata map[string][]mOKX.TypeKd

// 当前的分析结果
var TickerAnaly dbType.AnalyTickerType

// 当前币安的持仓
var BinancePositionList []binance.PositionType
