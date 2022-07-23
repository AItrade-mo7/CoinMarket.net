package ready

import (
	"time"

	"CoinMarket.net/server/okxApi/restApi/inst"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	mCycle.New(mCycle.Opt{
		Func:      GetInst,
		SleepTime: time.Hour * 4, // 每 4 时获取一次
	}).Start()
}

func GetInst() {
	inst.SPOT()
	time.Sleep(1 * time.Second) // 请求方法延迟 1 秒钟
	inst.SWAP()
	time.Sleep(1 * time.Second) // 请求方法延迟 1 秒钟
}
