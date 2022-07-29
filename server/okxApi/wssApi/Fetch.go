package wssApi

import (
	"github.com/EasyGolang/goTools/mFetch"
)

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
