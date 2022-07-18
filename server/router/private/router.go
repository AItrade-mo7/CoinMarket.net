package private

import (
	"CoinMarket.net/server/router/api/account"
	"github.com/gofiber/fiber/v2"
)

/*

/api/private

*/

func Router(router fiber.Router) {
	r := router.Group("/private", MiddleWare)

	r.Get("/get_user_info", account.GetUserInfo)
}
