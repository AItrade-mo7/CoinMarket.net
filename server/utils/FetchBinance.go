package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
)

type FetchBinanceOpt struct {
	Path          string
	Data          map[string]any
	Method        string
	IsLocalJson   bool
	LocalJsonPath string // 本地的数据源
	Event         func(string, any)
}

var BaseUrlArr = []string{
	"https://api.binance.com",
	"https://api1.binance.com",
	"https://api2.binance.com",
	"https://api3.binance.com",
}

func FetchBinance(opt FetchBinanceOpt) (resData []byte, resErr error) {
	// 是否为本地模式
	if opt.IsLocalJson {
		isJsonPath := mPath.Exists(opt.LocalJsonPath)
		if isJsonPath {
			return os.ReadFile(opt.LocalJsonPath)
		} else {
			resErr = fmt.Errorf("LocalJsonPath")
			return
		}
	}

	if len(opt.Method) < 1 {
		opt.Method = "GET"
	}

	// 处理 Header 和 加密信息
	Method := strings.ToUpper(opt.Method)

	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: BaseUrlArr[2],
		Path:   opt.Path,
		Data:   mJson.ToJson(opt.Data),
		Event:  opt.Event,
	})

	if Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
