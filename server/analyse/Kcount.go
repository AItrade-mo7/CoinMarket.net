package analyse

import "CoinMarket.net/server/okxApi/okxInfo"

// 构造新的Kdata
func NewKdata(data okxInfo.CandleDataType) (kdata okxInfo.Kd) {
	kdata = okxInfo.Kd{}

	return kdata
}
