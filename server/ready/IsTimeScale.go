package ready

import (
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

func IsTimeScale(KTime int64) bool {
	nowTimeObj := mTime.MsToTime(KTime, "0")

	Minute := nowTimeObj.Minute()

	isIn := false
	timeScale := []int{0, 15, 30, 45}
	for _, val := range timeScale {
		if Minute-val == 0 {
			isIn = true
			break
		}
	}
	return isIn
}

func GetKdataTime(TickerVol []mOKX.TypeTicker) int64 {
	var TimeUnix int64

	if len(TickerVol) > 0 {
		TimeUnix = TickerVol[0].Ts
	}

	return TimeUnix + 6000
}
