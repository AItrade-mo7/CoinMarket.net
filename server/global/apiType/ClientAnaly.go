package apiType

type ClientAnalyType struct {
	MaxUP           string `bson:"MaxUP"`
	MaxUP_RosePer   string `bson:"MaxUP_RosePer"`
	MaxDown         string `bson:"MaxDown"`
	MaxDown_RosePer string `bson:"MaxDown_RosePer"`
	Unit            string `bson:"Unit"`
	TimeUnix        int64  `bson:"TimeUnix"`
	TimeStr         string `bson:"TimeStr"`
	TimeID          string `bson:"TimeID"`
}