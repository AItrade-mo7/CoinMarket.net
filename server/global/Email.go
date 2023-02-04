package global

import (
	"CoinMarket.net/server/global/config"
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
		Account:     config.Email.Account,
		Password:    config.Email.Password,
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
