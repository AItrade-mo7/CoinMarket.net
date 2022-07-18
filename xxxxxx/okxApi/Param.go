package okxApi

type BaseUrlType struct {
	Rest   string
	PubWss string
	PriWss string
}

/*
okxApi.BaseUrl.Rest
okxApi.BaseUrl.PubWss
okxApi.BaseUrl.PriWss
*/
var BaseUrl = BaseUrlType{
	Rest:   "https://www.okx.com",
	PubWss: "wss://ws.okx.com:8443/ws/v5/public",
	PriWss: "wss://ws.okx.com:8443/ws/v5/private",
}
