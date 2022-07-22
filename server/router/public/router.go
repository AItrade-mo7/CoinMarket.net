package public

import (
	"CoinMarket.net/server/router/api"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/

func Router(router fiber.Router) {
	r := router.Group("/public", MiddleWare)
	r.Post("/ping", api.Ping)
}
