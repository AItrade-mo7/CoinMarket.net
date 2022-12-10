package tickerAnaly

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTime"
)

type SingleType struct {
	List  []mOKX.TypeKd // list
	Info  mOKX.AnalySingleType
	Slice map[int]mOKX.AnalySliceType

	ResData []mOKX.AnalySliceType
}

func NewSingle(list []mOKX.TypeKd) (_this *SingleType) {
	_this = &SingleType{}

	if len(list) < config.KdataLen {
		global.LogErr("tickerAnaly.NewSingle list 长度不足", len(list), list[0].InstID, list[0].TimeStr)
		return
	}

	size := len(list)
	_this.List = make([]mOKX.TypeKd, size)
	copy(_this.List, list)
	_this.Info.InstID = list[0].InstID

	_this.Slice = make(map[int]mOKX.AnalySliceType)

	_this.SetInfo()
	AnalySliceList := []mOKX.AnalySliceType{}
	for _, item := range config.SliceHour {
		_this.Slice[item] = _this.SliceKdata(item)
		sliceInfo := _this.AnalySlice(item)
		AnalySliceList = append(AnalySliceList, sliceInfo)
	}

	_this.ResData = AnalySliceList

	return
}

// 设置大数据的起止时间
func (_this *SingleType) SetInfo() *SingleType {
	list := _this.List
	Len := len(_this.List)

	_this.Info.StartTimeStr = list[0].TimeStr
	_this.Info.StartTimeUnix = list[0].TimeUnix
	_this.Info.EndTimeStr = list[Len-1].TimeStr
	_this.Info.EndTimeUnix = list[Len-1].TimeUnix
	_this.Info.DiffHour = (_this.Info.EndTimeUnix - _this.Info.StartTimeUnix) / mTime.UnixTimeInt64.Hour

	return _this
}

// 对数据进行切片
func (_this *SingleType) SliceKdata(hour int) (resData mOKX.AnalySliceType) {
	resData = mOKX.AnalySliceType{}
	list := _this.List
	Len := len(_this.List)

	startItem := list[Len-1-hour]
	lastItem := list[Len-1]

	resData.StartTimeStr = startItem.TimeStr
	resData.StartTimeUnix = startItem.TimeUnix
	resData.EndTimeStr = lastItem.TimeStr
	resData.EndTimeUnix = lastItem.TimeUnix
	DiffHour := (resData.EndTimeUnix - resData.StartTimeUnix) / mTime.UnixTimeInt64.Hour
	resData.DiffHour = int(DiffHour)

	resData.Len = hour + 1

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
	VolumeHour := make(map[string][]string)

	U_shade := []string{}
	D_shade := []string{}
	HLPer := []string{}
	for _, item := range list {

		Volume = mCount.Add(Volume, item.Vol)

		U_shade = append(U_shade, item.U_shade)
		D_shade = append(D_shade, item.D_shade)
		HLPer = append(HLPer, item.HLPer)

		TimeKey := mTime.MsToTime(item.TimeUnix, "0").Format("2006-01-02_15")
		VolumeHour[TimeKey] = append(VolumeHour[TimeKey], item.Vol)
	}

	VolumeHourArr := []string{}
	for _, l := range VolumeHour {
		Vol := "0"
		for _, v := range l {
			Vol = mCount.Add(Vol, v)
		}
		VolumeHourArr = append(VolumeHourArr, Vol)
	}

	Sort_H := mOKX.Sort_H(list)         // 最高价排序 高 - 低
	Sort_L := mOKX.Sort_L(list)         // 最低价排序 高 - 低
	Sort_HLPer := mOKX.Sort_HLPer(list) // 振幅排序 高 - 低

	slice.Volume = Volume
	slice.RosePer = mCount.RoseCent(lastElm.C, firstElm.C) // 最后一个的 C - 一开始的 CBas
	slice.H = Sort_H[0].H
	slice.L = Sort_L[len(Sort_H)-1].L

	U_shadeAvg := mCount.Average(U_shade)
	slice.U_shadeAvg = mCount.Cent(U_shadeAvg, 3)

	D_shadeAvg := mCount.Average(D_shade)
	slice.D_shadeAvg = mCount.Cent(D_shadeAvg, 3)

	slice.HLPerMax = Sort_HLPer[0].HLPer

	HLPerAvg := mCount.Average(HLPer)
	slice.HLPerAvg = mCount.Cent(HLPerAvg, 3)

	VolumeAvg := mCount.Average(VolumeHourArr)
	slice.VolumeAvg = mCount.Cent(VolumeAvg, 3)

	_this.Slice[Index] = slice

	return slice
}
