package config

import "github.com/EasyGolang/goTools/mStr"

// var SliceHour = []int{3, 5, 8, 12, 16, 20, 24, 28, 32, 36, 40}
// var SliceHour = []int{3, 5, 8, 12, 16, 20, 24}

var SliceHour = []int{3, 6, 9, 12}

var KdataLen = 100

/*
 [1m/3m/5m/15m/30m/1H/2H/4H]
*/

// mOKX.KdataBarOpt 的 key 值
var KdataBarOpt = []string{
	"1m",
	"3m",
	"5m",
	"15m",
	"30m",
	"1h",
	"2h",
	"4h",
}

var KdataBar = KdataBarOpt[5]

// 计价的锚定货币
var Unit = "USDT"

var SPOT_suffix = mStr.Join("-", Unit)

var SWAP_suffix = mStr.Join(SPOT_suffix, "-SWAP")
