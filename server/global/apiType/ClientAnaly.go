package apiType

import "github.com/EasyGolang/goTools/mOKX"

type ClientAnalyType struct {
	MillionCoin     []mOKX.AnalySliceType `bson:"MillionCoin"`
	Version         int                   `bson:"Version"`
	MaxUP           string                `bson:"MaxUP"`
	MaxUP_RosePer   string                `bson:"MaxUP_RosePer"`
	MaxDown         string                `bson:"MaxDown"`
	MaxDown_RosePer string                `bson:"MaxDown_RosePer"`
	Unit            string                `bson:"Unit"`
	WholeDir        int                   `bson:"WholeDir"`
	DirIndex        int                   `bson:"DirIndex"`
	TimeUnix        int64                 `bson:"TimeUnix"`
	TimeStr         string                `bson:"TimeStr"`
	TimeID          string                `bson:"TimeID"`
}
