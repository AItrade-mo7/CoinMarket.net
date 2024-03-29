package public

import (
	"CoinMarket.net/server/router/api"
	"github.com/gofiber/fiber/v2"
)

/*
/CoinMarket/public
*/

func Router(router fiber.Router) {
	r := router.Group("/public", MiddleWare)

	r.Post("/GetNowTickerAnaly", api.GetNowTickerAnaly)
	r.Post("/GetNowKdata", api.GetNowKdata)
	r.Post("/GetCoinHistory", api.GetCoinHistory)
	r.Post("/GetAnalyList", api.GetAnalyList)
	r.Post("/GetAnalyDetail", api.GetAnalyDetail)
	r.Post("/GetInstAll", api.GetInstAll)
	r.Post("/GetNowTrend", api.GetNowTrend)
}
