package ready

import (
	"github.com/EasyGolang/goTools/mTime"
)

func IsOKXDataTimeScale(KTime int64) bool {
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

func IsMongoDBTimeScale(KTime int64) bool {
	nowTimeObj := mTime.MsToTime(KTime, "0")

	Hour := nowTimeObj.Hour()

	isIn := false
	timeScale := []int{5, 13, 21}
	for _, val := range timeScale {
		if Hour-val == 0 {
			isIn = true
			break
		}
	}
	return isIn
}
