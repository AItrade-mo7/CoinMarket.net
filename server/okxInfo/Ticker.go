package okxInfo

type BinanceTickerType struct {
	Symbol             string `json:"symbol"`
	InstID             string `json:"InstID"`
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

var BinanceTickerList []BinanceTickerType // 币安的Ticker 排行

type OKXTickerType struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
}

var OKXTickerList []OKXTickerType // okx的Ticker

type TickerType struct {
	InstID         string `json:"InstID"`         // 产品ID
	CcyName        string `json:"CcyName"`        // 币种名称
	Last           string `json:"Last"`           // 最新成交价
	Open24H        string `json:"Open24H"`        // 24小时开盘价
	High24H        string `json:"High24H"`        // 最高价
	Low24H         string `json:"Low24H"`         // 最低价
	OKXVol24H      string `json:"OKXVol24H"`      // OKX 24小时成交量 USDT 数量
	BinanceVol24H  string `json:"BinanceVol24H"`  // 24 小时成交 USDT 数量
	U_R24          string `json:"U_R24"`          // 涨幅 = (最新价-开盘价)/开盘价 =
	U_RIdx         int    `json:"U_RIdx"`         // 涨幅 = (最新价-开盘价)/开盘价 =
	Volume         string `json:"Volume"`         // 成交量总和
	VolIdx         int    `json:"VolIdx"`         // 成交量排名
	OkxVolRose     string `json:"OkxVolRose"`     // 欧意占比总交易量
	BinanceVolRose string `json:"BinanceVolRose"` // 币安占比总交易量
	Ts             string `json:"Ts"`
}

var (
	TickerList  []TickerType
	TickerU_R24 []TickerType
)

// 市场分析
type WholeTickerAnalyseType struct {
	UPIndex  string     `json:"UPIndex"`  // 上涨指数
	UDAvg    string     `json:"UDAvg"`    // 综合涨幅均值
	UPLe     int        `json:"UPLe"`     // 上涨趋势
	UDLe     int        `json:"UDLe"`     // 上涨强度
	DirIndex int        `json:"DirIndex"` // 当前市场情况  -1 下跌   0 震荡   1 上涨
	MaxUP    TickerType `json:"MaxUP"`    // 最大涨幅币种
	MaxDown  TickerType `json:"MaxDown"`  // 最大跌幅币种
	Ts       int64      `json:"Ts"`       // 生成时间
}

var WholeTickerAnalyse WholeTickerAnalyseType
