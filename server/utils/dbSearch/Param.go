package dbSearch

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	SortType  map[string]int
	MatchType map[string]any
	QueryType map[string]any
	TimeType  [2]int64
)

type PagingType struct {
	List           []any     `bson:"List"`
	Total          int64     `bson:"Total"`
	Current        int64     `bson:"Current"`
	Size           int64     `bson:"Size"`
	Sort           SortType  `bson:"Sort"`           // key:value  -1 倒序 1：正序
	Match          MatchType `bson:"Match"`          // 匹配
	Query          QueryType `bson:"Query"`          // 查询
	CreateTimeUnix TimeType  `bson:"CreateTimeUnix"` // 创建时间查询
	TimeUnix       TimeType  `bson:"TimeUnix"`       // 更新时间查询
}

type FindParam struct {
	Size           int64     `bson:"Size"`           // 每页多少条
	Current        int64     `bson:"Current"`        // 当前页码 0 为第一页
	Sort           SortType  `bson:"Sort"`           // 排序
	Match          MatchType `bson:"Match"`          // 匹配
	Query          QueryType `bson:"Query"`          // 查询
	CreateTimeUnix TimeType  `bson:"CreateTimeUnix"` // 创建时间查询
	TimeUnix       TimeType  `bson:"TimeUnix"`       // 更新时间查询
	Type           string    `bson:"Type"`           // Serve  && Client  全面数据和简陋数据
}

type CurOpt struct {
	Param  FindParam
	DB     *mMongo.DB
	Total  int64
	Cursor *mongo.Cursor
}

func GetCursor(opt CurOpt) (resCur *CurOpt, resErr error) {
	resCur = &opt
	resErr = nil

	json := opt.Param

	if len(json.Sort) < 1 {
		sort := make(SortType)
		sort["CreateTime"] = -1
		json.Sort = sort
	}

	if len(json.Sort) > 1 {
		resErr = fmt.Errorf("%+v 参数数量太多", json.Sort)
		return
	}

	db := opt.DB

	// 构建搜索参数
	FK := bson.D{}

	// 构建匹配参数
	for key, val := range json.Match {

		var_arr := strings.Split(mStr.ToStr(val), `,`)

		for _, v := range var_arr {
			rgxStr := mStr.Join("^.*", mStr.ToStr(v), ".*$")
			FK = append(FK, bson.E{
				Key: key,
				Value: bson.D{
					{
						Key:   "$regex",
						Value: primitive.Regex{Pattern: rgxStr, Options: "i"},
					},
				},
			})
		}

	}
	// 构建查询参数
	for key, val := range json.Query {
		FK = append(FK, bson.E{
			Key:   key,
			Value: val,
		})
	}

	// 构建时间范围查询
	if (json.CreateTimeUnix[0] + json.CreateTimeUnix[1]) > 946656000000 {
		FK = append(FK, bson.E{
			Key: "CreateTimeUnix",
			Value: bson.D{
				{
					Key:   "$gte", // 大于或等于
					Value: json.CreateTimeUnix[0],
				}, {
					Key:   "$lte", // 小于或等于
					Value: json.CreateTimeUnix[1],
				},
			},
		})
	}

	opt.Param = json

	// 查询总条目
	total, err := db.Table.CountDocuments(db.Ctx, FK)
	if err != nil {
		db.Close()
		resErr = fmt.Errorf("读取总条目失败 %+v", err)
		return
	}
	resCur.Total = total

	findOpt := FindOpt(json)

	cur, err := db.Table.Find(db.Ctx, FK, findOpt)
	if err != nil {
		db.Close()
		resErr = fmt.Errorf("数据读取失败 %+v", err)
		return
	}
	resCur.Cursor = cur

	return
}

func FindOpt(json FindParam) *options.FindOptions {
	findOpt := options.Find()
	findOpt.SetSort(json.Sort)
	findOpt.SetSkip(json.Current * json.Size)
	findOpt.SetLimit(json.Size)
	findOpt.SetAllowDiskUse(true)

	return findOpt
}

func (obj *CurOpt) GenerateData(list []any) PagingType {
	json := obj.Param

	var returnData PagingType
	returnData.List = list
	returnData.Current = json.Current
	returnData.Total = obj.Total
	returnData.Size = json.Size
	returnData.Sort = json.Sort
	returnData.Match = json.Match
	returnData.Query = json.Query
	returnData.CreateTimeUnix = json.CreateTimeUnix
	returnData.TimeUnix = json.TimeUnix

	return returnData
}
