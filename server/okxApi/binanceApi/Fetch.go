package binanceApi

import (
	"strings"

	"github.com/EasyGolang/goTools/mFetch"
)

/*
	resData, err := restApi.Fetch(restApi.FetchOpt{
		Path: "/api/v5/account/balance",
		Data: map[string]any{
			"xxxx": "xxxx",
		},
		Method: "get",
		Event: func(s string, a any) {
			fmt.Println("Event", s, a)
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(mStr.ToStr(resData))

*/

type FetchOpt struct {
	Path   string
	Data   map[string]any
	Method string
	Event  func(string, any)
}

func Fetch(opt FetchOpt) (resData []byte, resErr error) {
	if len(opt.Method) < 1 {
		opt.Method = "GET"
	}

	// 处理 Header 和 加密信息
	Method := strings.ToUpper(opt.Method)

	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "https://api2.binance.com",
		Path:   opt.Path,
		Data:   opt.Data,
		Event:  opt.Event,
	})

	if Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
