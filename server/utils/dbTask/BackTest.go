package dbTask

import "fmt"

type BackTestOpt struct {
	StartTime string
	EndTime   string
}

type BackTestObj struct {
	StartTime int64
	EndTime   int64
}

func BackTest(opt BackTestOpt) *BackTestObj {
	obj := BackTestObj{}

	fmt.Println(opt.StartTime)
	fmt.Println(opt.EndTime)

	// 回填数据
	obj.FillKdata()

	return &obj
}

func (obj *BackTestObj) FillKdata() {
}
