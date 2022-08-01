package tickerAnalyse

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type SliceType struct {
	List  []okxInfo.Kd
	Slice okxInfo.AnalyseSliceType
}

type SingleType struct {
	List   []okxInfo.Kd // list
	Single okxInfo.AnalyseSingleType
	Slice  map[int]SliceType
}

func NewSingle(list []okxInfo.Kd) *SingleType {
	if len(list) != 300 { // 数组不为300条的一概不理睬
		return nil
	}
	_this := &SingleType{}
	size := len(list)
	_this.List = make([]okxInfo.Kd, size)
	copy(_this.List, list)
	_this.Single.InstID = list[0].InstID
	_this.Slice = make(map[int]SliceType)

	_this.SetTime()
	SliceHour := []int{1, 2, 4, 8, 12, 16, 24}
	for _, item := range SliceHour {
		_this.Slice[item] = _this.SliceKdata(item)
	}

	return _this
}

// 设置大数据的起止时间
func (_this *SingleType) SetTime() *SingleType {
	list := _this.List
	Len := len(_this.List)

	_this.Single.StartTime = list[0].Time
	_this.Single.StartTimeUnix = list[0].TimeUnix
	_this.Single.EndTime = list[Len-1].Time
	_this.Single.EndTimeUnix = list[Len-1].TimeUnix
	_this.Single.DiffHour = (_this.Single.EndTimeUnix - _this.Single.StartTimeUnix) / mTime.UnixTimeInt64.Hour

	return _this
}

// 对数据进行切片
func (_this *SingleType) SliceKdata(hour int) (resData SliceType) {
	resData = SliceType{}
	list := _this.List
	Len := len(_this.List)

	backward := int64(hour)
	nowTimeUnix := list[Len-1].TimeUnix
	tarTime := mTime.MsToTime(nowTimeUnix, mStr.Join("-", backward*mTime.UnixTimeInt64.Hour))

	tarTimeUnix := mTime.ToUnixMsec(tarTime)
	diffM := int64(tarTime.Minute()) * (mTime.UnixTimeInt64.Minute)

	startTimeUnix := tarTimeUnix - diffM

	for _, item := range list {
		if item.TimeUnix >= startTimeUnix {
			resData.List = append(resData.List, item)
		}
	}

	cList := resData.List
	cLen := len(resData.List)

	resData.Slice.StartTime = cList[0].Time
	resData.Slice.StartTimeUnix = cList[0].TimeUnix
	resData.Slice.EndTime = cList[cLen-1].Time
	resData.Slice.EndTimeUnix = cList[cLen-1].TimeUnix
	DiffHour := (resData.Slice.EndTimeUnix - resData.Slice.StartTimeUnix) / mTime.UnixTimeInt64.Hour
	resData.Slice.DiffHour = int(DiffHour)

	return
}

// 对切片数据进行分析
/*
最高价、最低价、震动均值、首尾价差、
*/
