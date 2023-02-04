package ready

import (
	"github.com/EasyGolang/goTools/mTime"
)

func IsOKXDataTimeScale(KTime int64) bool {
	nowTimeObj := mTime.MsToTime(KTime, "0")

	Minute := nowTimeObj.Minute()

	isIn := false
	timeScale := []int{1, 16, 31, 46}
	for _, val := range timeScale {
		if Minute-val == 0 {
			isIn = true
			break
		}
	}
	return isIn
}

func IsRestartShellTimeScale(KTime int64) bool {
	nowTimeObj := mTime.MsToTime(KTime, "0")
	Hour := nowTimeObj.Hour()

	isIn := false
	timeScale := []int{3} // 小时数为 3 的时候 执行一次重启数据库操作
	for _, val := range timeScale {
		if Hour-val == 0 {
			isIn = true
			break
		}
	}

	return isIn
}
