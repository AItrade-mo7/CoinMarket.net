package public

import (
	"CoinMarket.net/server/router/api"
	"CoinMarket.net/server/router/api/account"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/

func Router(router fiber.Router) {
	r := router.Group("/public", MiddleWare)

	r.Post("/login", account.Login)
	r.Post("/upload", api.Upload)
}
