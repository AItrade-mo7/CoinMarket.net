package config

var SysEnv struct {
	MongoAddress  string
	MongoPassword string
	MongoUserName string
}

func LoadSysEnv() {
	SysEnv.MongoAddress = "trade.mo7.cc:17017"
	SysEnv.MongoPassword = "asdasd55555"
	SysEnv.MongoUserName = "mo7"
}

var AppInfo struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
	Port    string `json:"Port"`
}
