package dbType

import "github.com/EasyGolang/goTools/mEncrypt"

type AccountTable struct {
	Email    string `bson:"Email"`    // 用户 Email
	UserID   string `bson:"UserID"`   // 用户 ID
	Avatar   string `bson:"Avatar"`   // 用户头像
	NickName string `bson:"NickName"` // 用户昵称
	Password string `bson:"Password"` // 用户密码
}

var UserInfo = AccountTable{
	NickName: "墨七",
	Avatar:   "//file.mo7.cc/static/lxh_png/lxh.png",
	UserID:   "a8c776fca8a046ee8d5c43aa374849c4",
	Email:    "meichangliang@mo7.cc",
	Password: mEncrypt.MD5("123123"),
}
