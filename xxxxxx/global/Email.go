package global

import "github.com/EasyGolang/goTools/mEmail"

type EmailOpt struct {
	To       []string
	Subject  string
	Template string
	SendData any
}

func Email(opt EmailOpt) *mEmail.EmailInfo {
	emailObj := mEmail.New(mEmail.Opt{
		Account:     UserEnv.Email.Account,
		Password:    UserEnv.Email.Password,
		To:          opt.To,
		From:        "Hunter数据中心",
		Subject:     opt.Subject,
		Port:        "587",
		Host:        "smtp.feishu.cn",
		TemplateStr: opt.Template,
		SendData:    opt.SendData,
	})

	return emailObj
}
