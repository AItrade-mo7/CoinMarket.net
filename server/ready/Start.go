package ready

import (
	"time"

	"CoinMarket.net/server/okxApi/restApi/inst"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	mCycle.New(mCycle.Opt{
		Func:      inst.Start,
		SleepTime: time.Hour * 4, // 每 4 时获取一次
	}).Start()
}
