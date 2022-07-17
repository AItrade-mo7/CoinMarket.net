package okxInfo

type Kd struct {
	InstID   string `bson:"InstID"`   // 持仓币种
	TimeUnix string `bson:"TimeUnix"` // 毫秒时间戳
	O        string `bson:"O"`        // 开盘
	H        string `bson:"H"`        // 最高
	L        string `bson:"L"`        // 最低
	C        string `bson:"C"`        // 收盘价格
	DataType string `bson:"DataType"` // history 或者 WSS
	KRunType string `bson:"KRunType"` // K 线的类型 tradeKRun  &&  dirKRun
	Dir      int    `bson:"Dir"`      // 方向 (收盘-开盘) ，1：涨 & -1：跌 & 0：横盘
	Center   string `bson:"Center"`   // 实体中心价 (开盘+收盘+最高+最低) / 4
	HLPer    string `bson:"HLPer"`    // (最高-最低)/最低 * 100%
	SolidPer string `bson:"SolidPer"` // 实体的百分点(收盘-开盘)/开盘
	RosePer  string `bson:"RosePer"`  // 涨幅

	C_dir   int    `bson:"C_dir"`   // 中心点方向 (当前中心点-前中心点) 1：涨 & -1：跌 & 0：横盘
	U_shade string `bson:"U_shade"` // 上影线
	D_shade string `bson:"D_shade"` // 下影线

	Ema5  string `bson:"Ema5"`
	Ema15 string `bson:"Ema15"`
	Ema60 string `bson:"Ema60"`
	Ma5   string `bson:"Ma5"`
	Ma15  string `bson:"Ma15"`
	Ma60  string `bson:"Ma60"`

	Sar    string `bson:"Sar"`    // Sar 指标
	SarDir int    `bson:"SarDir"` // SarDir 指标方向 1  -1
}
