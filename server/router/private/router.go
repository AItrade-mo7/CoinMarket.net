package private

import (
	"CoinMarket.net/server/router/api"
	"github.com/gofiber/fiber/v2"
)

/*

/api/private

*/

func Router(router fiber.Router) {
	r := router.Group("/private", MiddleWare)

	r.Post("/ping", api.Ping)
}
