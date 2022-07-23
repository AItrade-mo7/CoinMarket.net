package wssApi

import (
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
type LoginArgsType struct {
	APIKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}
type LoginType struct {
	Op   string          `json:"op"`
	Args []LoginArgsType `json:"args"`
}

type FetchOpt struct {
	Type  int
	Event func(string, any)
}

func New(opt FetchOpt) (_this *mFetch.Wss) {
	WssOpt := mFetch.WssOpt{}
	WssOpt.Event = opt.Event
	if opt.Type == 0 {
		WssOpt.Url = "wss://ws.okx.com:8443/ws/v5/public"
	}
	if opt.Type == 1 {
		WssOpt.Url = "wss://ws.okx.com:8443/ws/v5/private"
	}
	_this = mFetch.NewWss(WssOpt)

	return
}
