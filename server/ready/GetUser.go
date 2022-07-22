package ready

import (
	"fmt"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/utils/dbUser"
)

func GetUserInfo() {
	UserID := config.AppEnv.UserID

	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		UserDB.DB.Close()
		errStr := fmt.Errorf("用户数据读取错误 %+v", err)
		global.LogErr(errStr)
		return
	}
}
