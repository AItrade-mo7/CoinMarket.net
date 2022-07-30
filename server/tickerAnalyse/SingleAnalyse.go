package tickerAnalyse

import (
	"time"

	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

/*
单个币种历史数据分析

需要分析的部分：
近1小时上涨情况
近2小时上涨情况
近3小时上涨情况
近4小时上涨情况
近5小时上涨情况

榜单整体上涨情况
*/

type NewSingleType struct {
	List          []okxInfo.Kd `json:"List"`      // list
	InstID        string       `json:"InstID"`    // InstID
	StartTime     time.Time    `json:"StartTime"` // 开始时间
	StartTimeUnix int64        `json:"StartTimeUnix"`
	EndTime       time.Time    `json:"EndTime"` // 结束时间
	EndTimeUnix   int64        `json:"EndTimeUnix"`
	DiffHour      int64        `json:"DiffHour"` // 总时长
	List1         []okxInfo.Kd `json:"List1"`    // 1 小时切片
	List2         []okxInfo.Kd `json:"List2"`    // 2 小时切片
	List4         []okxInfo.Kd `json:"List4"`    // 4 小时切片
	List8         []okxInfo.Kd `json:"List8"`    // 8 小时切片
	List12        []okxInfo.Kd `json:"List12"`   // 12 小时切片
	List24        []okxInfo.Kd `json:"List24"`   // 24 小时切片
}

func NewSingle(list []okxInfo.Kd) *NewSingleType {
	if len(list) != 300 { // 数组不为300条的一概不理睬
		return nil
	}
	_this := &NewSingleType{}
	size := len(list)
	_this.List = make([]okxInfo.Kd, size)
	copy(_this.List, list)
	_this.InstID = list[0].InstID

	_this.SetTime()
	_this.List1 = _this.SliceList(1)
	_this.List2 = _this.SliceList(2)
	_this.List4 = _this.SliceList(4)
	_this.List8 = _this.SliceList(8)
	_this.List12 = _this.SliceList(12)
	_this.List24 = _this.SliceList(24)

	return _this
}

// 设置起止时间
func (_this *NewSingleType) SetTime() *NewSingleType {
	list := _this.List
	Len := len(_this.List)

	_this.StartTime = list[0].Time
	_this.StartTimeUnix = list[0].TimeUnix
	_this.EndTime = list[Len-1].Time
	_this.EndTimeUnix = list[Len-1].TimeUnix
	_this.DiffHour = (_this.EndTimeUnix - _this.StartTimeUnix) / mTime.UnixTimeInt64.Hour

	return _this
}

func (_this *NewSingleType) SliceList(hour int64) (resData []okxInfo.Kd) {
	resData = []okxInfo.Kd{}
	list := _this.List
	Len := len(_this.List)

	backward := hour
	nowTimeUnix := list[Len-1].TimeUnix
	tarTime := mTime.MsToTime(nowTimeUnix, mStr.Join("-", backward*mTime.UnixTimeInt64.Hour))

	tarTimeUnix := mTime.ToUnixMsec(tarTime)
	diffM := int64(tarTime.Minute()) * (mTime.UnixTimeInt64.Minute)

	startTimeUnix := tarTimeUnix - diffM

	for _, item := range list {
		if item.TimeUnix >= startTimeUnix {
			resData = append(resData, item)
		}
	}

	return
}
