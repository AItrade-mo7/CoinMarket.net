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

	r.Post("/Tickers", api.Tickers)
	r.Post("/Inst", api.Inst)
	r.Post("/GetAnalyKdata", api.GetAnalyKdata)
}
