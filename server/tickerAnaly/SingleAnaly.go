package tickerAnaly

import (
	"CoinMarket.net/server/okxInfo"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type SingleType struct {
	List  []mOKX.TypeKd // list
	Info  mOKX.AnalySingleType
	Slice map[int]mOKX.AnalySliceType
}

func NewSingle(list []mOKX.TypeKd) *SingleType {
	if len(list) != 300 { // 数组不为300条的一概不理睬
		return nil
	}
	_this := &SingleType{}
	size := len(list)
	_this.List = make([]mOKX.TypeKd, size)
	copy(_this.List, list)
	_this.Info.InstID = list[0].InstID
	_this.Slice = make(map[int]mOKX.AnalySliceType)

	_this.SetTime()
	SliceHour := []int{1, 2, 4, 8, 12, 16, 24}
	AnalySliceList := []mOKX.AnalySliceType{}
	for _, item := range SliceHour {
		_this.Slice[item] = _this.SliceKdata(item)
		sliceInfo := _this.AnalySlice(item)
		AnalySliceList = append(AnalySliceList, sliceInfo)
	}

	okxInfo.TickerAnalySingle[_this.Info.InstID] = AnalySliceList
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
func (_this *SingleType) SliceKdata(hour int) (resData mOKX.AnalySliceType) {
	resData = mOKX.AnalySliceType{}
	list := _this.List
	Len := len(_this.List)

	// 切片数组
	cList := []mOKX.TypeKd{}

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

// 获取数组
func (_this *SingleType) GetSliceList(Index int) []mOKX.TypeKd {
	Slice := _this.Slice[Index]
	AllLen := len(_this.List)
	Len := Slice.Len
	List := _this.List[AllLen-Len : AllLen]
	size := len(List)
	reList := make([]mOKX.TypeKd, size)
	copy(reList, List)
	return reList
}

func (_this *SingleType) AnalySlice(Index int) mOKX.AnalySliceType {
	slice := _this.Slice[Index]
	list := _this.GetSliceList(Index)
	slice.InstID = list[0].InstID

	firstElm := list[0]
	lastElm := list[len(list)-1]

	Volume := "0" // 成交量总和
	VolumeArr := []string{}
	U_shade := []string{}
	D_shade := []string{}
	HLPer := []string{}
	for _, item := range list {
		Volume = mCount.Add(Volume, item.VolCcy)
		U_shade = append(U_shade, item.U_shade)
		D_shade = append(D_shade, item.D_shade)
		HLPer = append(HLPer, item.HLPer)
		VolumeArr = append(VolumeArr, item.VolCcy)
	}
	Sort_H := mOKX.Sort_H(list)         // 最高价排序 高 - 低
	Sort_L := mOKX.Sort_L(list)         // 最低价排序 高 - 低
	Sort_HLPer := mOKX.Sort_HLPer(list) // 最低价排序 高 - 低

	slice.Volume = Volume
	slice.RosePer = mCount.RoseCent(lastElm.C, firstElm.O) // 最后一个的收盘价 - 一开始的开盘价
	slice.H = Sort_H[0].H
	slice.L = Sort_L[len(Sort_H)-1].L

	U_shadeAvg := mCount.Average(U_shade)
	slice.U_shadeAvg = mCount.Cent(U_shadeAvg, 3)

	D_shadeAvg := mCount.Average(D_shade)
	slice.D_shadeAvg = mCount.Cent(D_shadeAvg, 3)

	slice.HLPerMax = Sort_HLPer[0].HLPer

	HLPerAvg := mCount.Average(HLPer)
	slice.HLPerAvg = mCount.Cent(HLPerAvg, 3)

	VolumeAvg := mCount.Average(VolumeArr)
	slice.VolumeAvg = mCount.Cent(VolumeAvg, 3)

	_this.Slice[Index] = slice

	return slice
}
