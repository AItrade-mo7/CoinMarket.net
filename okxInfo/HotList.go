package okxInfo

type HotInfo struct {
	InstType  string `bson:"InstType"`
	InstID    string `bson:"InstID"`
	Last      string `bson:"Last"` // 最新成交价
	LastSz    string `bson:"LastSz"`
	AskPx     string `bson:"AskPx"`
	AskSz     string `bson:"AskSz"` // 卖一价的挂单数数量
	BidPx     string `bson:"BidPx"`
	BidSz     string `bson:"BidSz"`     // 买一价的挂单数量
	Open24H   string `bson:"Open24H"`   // 开盘价
	High24H   string `bson:"High24H"`   // 最高价
	Low24H    string `bson:"Low24H"`    // 最低价
	VolCcy24H string `bson:"VolCcy24H"` // 24小时成交量 以币为单位
	Vol24H    string `bson:"Vol24H"`
	Ts        string `bson:"Ts"`
	SodUtc0   string `bson:"SodUtc0"`
	SodUtc8   string `bson:"SodUtc8"`

	// 自定义部分
	CcyName   string       `bson:"CcyName"`   // 币种名称
	U_R24     string       `bson:"U_R24"`     // 24H涨幅
	Amount    string       `bson:"Amount"`    // 成交额 = 平均价格  * 成交量
	Average24 string       `bson:"Average24"` // 平均价格 = 24h 开盘价 + 最高价格 + 最低价格 + 最新成交价 算数平均数  =  单位 USDT
	SWAP      InstInfoType `bson:"SWAP"`      // 该排行币种对应的 合约 产品数据
	SPOT      InstInfoType `bson:"SPOT"`      // 该排行榜单对应的 现货 产品数据
}

// 成交额排序
var Amount24Hot []HotInfo

// 涨跌幅排序
var U_R24Hot []HotInfo

// 涨跌幅绝对值
var U_R24AbsHot []HotInfo
