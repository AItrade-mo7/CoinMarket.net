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
	InstType  string `json:"InstType"`
	InstID    string `json:"InstID"`
	Last      string `json:"Last"` // 最新成交价
	LastSz    string `json:"LastSz"`
	AskPx     string `json:"AskPx"`
	AskSz     string `json:"AskSz"`
	BidPx     string `json:"BidPx"`
	BidSz     string `json:"BidSz"`
	Open24H   string `json:"Open24H"`   // 24小时开盘价
	High24H   string `json:"High24H"`   // 最高价
	Low24H    string `json:"Low24H"`    // 最低价
	VolCcy24H string `json:"VolCcy24H"` // 24小时成交量 USDT 数量
	// Binance 数据
	QuoteVolume string `json:"QuoteVolume"` // 24 小时成交 USDT 数量
	// 自定义数据
	U_R24   string `json:"U_R24"`   // 涨幅 = (最新价-开盘价)/开盘价 =
	CcyName string `json:"CcyName"` // 币种名称
	Volume  string `json:"Volume"`  // 成交量总和
}

var (
	TickerList  []TickerType
	TickerU_R24 []TickerType
)
