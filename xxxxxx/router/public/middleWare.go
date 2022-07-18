package public

import (
	"github.com/gin-gonic/gin"
)

func MiddleWare(c *gin.Context) {
	c.Writer.Header().Del("Data-Type")
	c.Header("Data-Type", "CoinMarket.net/public")

	c.Next()
}
