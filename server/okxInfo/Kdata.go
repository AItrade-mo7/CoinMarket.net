package okxInfo

import "github.com/EasyGolang/goTools/mOKX"

// 榜单币种 近 300 条数据 15 分钟间隔 共 75 小时
var MarketKdata = map[string][]mOKX.Kd{}

var TickerAnalyseSingle = map[string]mOKX.AnalyseSingleType{}
