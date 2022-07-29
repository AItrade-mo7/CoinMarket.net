package restApi

import (
	"strings"

	"github.com/EasyGolang/goTools/mFetch"
)

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
		Origin: "https://www.okx.com",
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
