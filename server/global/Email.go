package global

import (
	"github.com/EasyGolang/goTools/mEmail"
)

type EmailOpt struct {
	To       []string
	Subject  string
	Template string
	SendData any
}

func Email(opt EmailOpt) *mEmail.EmailInfo {
	emailObj := mEmail.New(mEmail.Opt{
		Account:     "trade@mo7.cc",
		Password:    "Mcl931750",
		To:          opt.To,
		From:        "CoinMarket 信息搜集",
		Subject:     opt.Subject,
		Port:        "587",
		Host:        "smtp.feishu.cn",
		TemplateStr: opt.Template,
		SendData:    opt.SendData,
	})

	return emailObj
}
