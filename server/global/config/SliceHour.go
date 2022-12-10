package config

// var SliceHour = []int{3, 5, 8, 12, 16, 20, 24, 28, 32, 36, 40}
// var SliceHour = []int{3, 5, 8, 12, 16, 20, 24}
var SliceHour = []int{3, 5, 8}

var KdataLen = 100

/*
 [1m/3m/5m/15m/30m/1H/2H/4H]
*/

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
