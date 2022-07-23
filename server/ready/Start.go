package ready

import (
	"time"

	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/okxApi/okxInfo"
	"CoinMarket.net/server/okxApi/restApi/inst"
	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
)

func Start() {
	// 获取 OKX 交易产品信息
	mCycle.New(mCycle.Opt{
		Func:      inst.Start,
		SleepTime: time.Hour * 4, // 每 4 时获取一次
	}).Start()
	mFile.Write(config.Dir.JsonData+"/SPOT_inst.json", mJson.ToStr(okxInfo.SPOT_inst))
	mFile.Write(config.Dir.JsonData+"/SWAP_inst.json", mJson.ToStr(okxInfo.SWAP_inst))
}
