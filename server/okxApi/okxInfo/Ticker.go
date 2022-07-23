package okxInfo

type TickerType struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"` // 最新成交价
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`   // 24小时开盘价
	High24H   string `json:"high24h"`   // 最高价
	Low24H    string `json:"low24h"`    // 最低价
	Avg24     string `json:"avg24"`     // 24 平均价格
	Amount    string `json:"amount"`    // 24 小时成交额
	VolCcy24H string `json:"volCcy24h"` // 24小时成交量 以币为单位
	U_R24     string `json:"u_r24"`     // 涨幅 = (最新价-开盘价)/开盘价 =
	CcyName   string `json:"CcyName"`   // 币种名称
}

var TickerList []TickerType
