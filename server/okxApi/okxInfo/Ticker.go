package okxInfo

type BinanceTickerType struct {
	Symbol             string `json:"symbol"`
	InstID             string `json:"instId"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int    `json:"firstId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

var BinanceTickerList []BinanceTickerType // 只用作数据中转

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
	VolCcy24H string `json:"volCcy24h"` // 24小时成交量 USDT 数量
	// Binance 数据
	QuoteVolume string `json:"quoteVolume"` // 24 小时成交 USDT 数量
	// 自定义数据
	U_R24   string `json:"u_r24"`   // 涨幅 = (最新价-开盘价)/开盘价 =
	CcyName string `json:"CcyName"` // 币种名称
	Amount  string `json:"amount"`  // 成交量总和
}

var (
	TickerList  []TickerType
	TickerU_R24 []TickerType
)
