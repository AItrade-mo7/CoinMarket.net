package okxInfo

import "time"

type CandleDataType [7]string

type Kd struct {
	InstID   string    `json:"InstID"`   // 持仓币种
	TimeUnix string    `json:"TimeUnix"` // 毫秒时间戳
	Time     time.Time `json:"Time"`     // 时间
	O        string    `json:"O"`        // 开盘
	H        string    `json:"H"`        // 最高
	L        string    `json:"L"`        // 最低
	C        string    `json:"C"`        // 收盘价格
	Vol      string    `json:"Vol"`      // 交易货币的数量
	VolCcy   string    `json:"VolCcy"`   // 计价货币数量
	// 自定义计算
	Dir      int    `json:"Dir"`      // 方向 (收盘-开盘) ，1：涨 & -1：跌 & 0：横盘
	Center   string `json:"Center"`   // 实体中心价 (开盘+收盘+最高+最低) / 4
	HLPer    string `json:"HLPer"`    // (最高-最低)/最低 * 100%
	SolidPer string `json:"SolidPer"` // 实体的百分点(收盘-开盘)/开盘
	RosePer  string `json:"RosePer"`  // 涨幅
	C_dir    int    `json:"C_dir"`    // 中心点方向 (当前中心点-前中心点) 1：涨 & -1：跌 & 0：横盘
	U_shade  string `json:"U_shade"`  // 上影线
	D_shade  string `json:"D_shade"`  // 下影线
	Type     string `json:"Type"`     // 数据类型
}
