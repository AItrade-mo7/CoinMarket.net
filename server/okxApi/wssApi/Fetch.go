package wssApi

import (
	"github.com/EasyGolang/goTools/mFetch"
)

type LoginArgsType struct {
	APIKey     string `bson:"apiKey"`
	Passphrase string `bson:"passphrase"`
	Timestamp  string `bson:"timestamp"`
	Sign       string `bson:"sign"`
}
type LoginType struct {
	Op   string          `bson:"op"`
	Args []LoginArgsType `bson:"args"`
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
