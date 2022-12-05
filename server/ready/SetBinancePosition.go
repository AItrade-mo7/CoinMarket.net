package ready

import (
	"CoinMarket.net/server/global"
	"CoinMarket.net/server/okxApi/binanceApi"
	"github.com/EasyGolang/goTools/mJson"
)

func SetBinancePosition() {
	BinancePosition := binanceApi.GetAccount() // 存储到数据库 BinancePosition
	global.Run.Println("读取一次 币安持仓", mJson.ToStr(BinancePosition))

	mJson.Println(BinancePosition)
}
