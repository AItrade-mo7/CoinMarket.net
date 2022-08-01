package tickerAnalyse

import (
	"fmt"

	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type SingleType struct {
	List  []okxInfo.Kd // list
	Info  okxInfo.AnalyseSingleType
	Slice map[int]okxInfo.AnalyseSliceType
}

func NewSingle(list []okxInfo.Kd) *SingleType {
	if len(list) != 300 { // 数组不为300条的一概不理睬
		return nil
	}
	_this := &SingleType{}
	size := len(list)
	_this.List = make([]okxInfo.Kd, size)
	copy(_this.List, list)
	_this.Info.InstID = list[0].InstID
	_this.Slice = make(map[int]okxInfo.AnalyseSliceType)

	_this.SetTime()
	SliceHour := []int{1, 2, 4, 8, 12, 16, 24}
	for _, item := range SliceHour {
		_this.Slice[item] = _this.SliceKdata(item)
		_this.AnalyseSlice(item)
	}

	return _this
}

// 设置大数据的起止时间
func (_this *SingleType) SetTime() *SingleType {
	list := _this.List
	Len := len(_this.List)

	_this.Info.StartTime = list[0].Time
	_this.Info.StartTimeUnix = list[0].TimeUnix
	_this.Info.EndTime = list[Len-1].Time
	_this.Info.EndTimeUnix = list[Len-1].TimeUnix
	_this.Info.DiffHour = (_this.Info.EndTimeUnix - _this.Info.StartTimeUnix) / mTime.UnixTimeInt64.Hour

	return _this
}

// 对数据进行切片
func (_this *SingleType) SliceKdata(hour int) (resData okxInfo.AnalyseSliceType) {
	resData = okxInfo.AnalyseSliceType{}
	list := _this.List
	Len := len(_this.List)

	// 切片数组
	cList := []okxInfo.Kd{}

	backward := int64(hour)
	nowTimeUnix := list[Len-1].TimeUnix
	tarTime := mTime.MsToTime(nowTimeUnix, mStr.Join("-", backward*mTime.UnixTimeInt64.Hour))

	tarTimeUnix := mTime.ToUnixMsec(tarTime)
	diffM := int64(tarTime.Minute()) * (mTime.UnixTimeInt64.Minute)

	startTimeUnix := tarTimeUnix - diffM

	for _, item := range list {
		if item.TimeUnix >= startTimeUnix {
			cList = append(cList, item)
		}
	}

	cLen := len(cList)

	resData.StartTime = cList[0].Time
	resData.StartTimeUnix = cList[0].TimeUnix
	resData.EndTime = cList[cLen-1].Time
	resData.EndTimeUnix = cList[cLen-1].TimeUnix
	DiffHour := (resData.EndTimeUnix - resData.StartTimeUnix) / mTime.UnixTimeInt64.Hour
	resData.DiffHour = int(DiffHour)
	resData.Len = len(cList)

	return
}

// 对切片数据进行分析
/*
最高价、最低价、震动均值、首尾价差、
*/

func (_this *SingleType) AnalyseSlice(Index int) {
	slice := _this.Slice[Index]
	list := _this.GetSliceList(Index)
	mJson.Println(slice)

	fmt.Println(list[0].Time, len(list), list[len(list)-1].Time)
}

// 获取数组
func (_this *SingleType) GetSliceList(Index int) []okxInfo.Kd {
	Slice := _this.Slice[Index]
	AllLen := len(_this.List)
	Len := Slice.Len
	List := _this.List[AllLen-Len : AllLen]

	size := len(List)
	reList := make([]okxInfo.Kd, size)
	copy(reList, List)

	return reList
}
