package middleWare

import (
	"net/http"

	"CoinMarket.net/config"
	"CoinMarket.net/utils/ginResult"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gin-gonic/gin"
)

func Public(c *gin.Context) {
	// 添加访问头
	AddHeader(c)

	// 授权验证
	EncryptAuth(c)

	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Del("Data-Type")
	c.Header("Data-Type", "CoinMarket.net")
}

func EncryptAuth(c *gin.Context) {
	AuthEncrypt := c.Request.Header["Auth-Encrypt"]
	if len(AuthEncrypt) < 1 {
		c.JSON(http.StatusOK, ginResult.AuthErr.WithData("需要添加授权码"))
		c.Abort()
		return
	}

	shaStr := mEncrypt.Sha256(
		mStr.Join(c.Request.URL.Path, "mo7"),
		config.SecretKey)

	if AuthEncrypt[0] != shaStr {
		c.JSON(http.StatusOK, ginResult.AuthErr.WithData("授权验证错误"))
		c.Abort()
		return
	}
}
