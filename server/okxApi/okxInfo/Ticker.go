package okxInfo

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
	U_R24          string `json:"U_R24"`          // 涨幅 = (最新价-开盘价)/开盘价 =
	CcyName        string `json:"CcyName"`        // 币种名称
	Volume         string `json:"Volume"`         // 成交量总和
	OkxVolRose     string `json:"OkxVolRose"`     // 欧意占比总交易量
	BinanceVolRose string `json:"BinanceVolRose"` // 币安占比总交易量
}

var (
	TickerList  []TickerType
	TickerU_R24 []TickerType
)
