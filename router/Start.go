package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"CoinMarket.net/global"
	"CoinMarket.net/router/middleWare"
	"CoinMarket.net/router/public"

	"CoinMarket.net/config"

	"github.com/EasyGolang/goTools/mStr"
	"github.com/gin-gonic/gin"
)

var FilePathArr []string

func Start() {
	logFile, _ := os.Create(config.LogPath + "/WebServer-T" + time.Now().Format("06年1月02日15时") + ".log")

	gin.DefaultWriter = io.MultiWriter(logFile)

	router := gin.Default()
	router.Use(
		middleWare.Public,
		middleWare.RateLimitMiddleware(time.Second, 100, 100),
	)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "欢迎访问 CoinMarket.net 服务")
	})

	// public
	public_g := router.Group("/public")
	public_g.Use(public.MiddleWare)
	{
		public_g.GET("/wss", public.WsServer)
		public.Router(public_g)
	}

	port := global.UserEnv.Port

	logStr := mStr.Join(`启动 CoinMarket.net 服务:  http://localhost:`, port)

	fmt.Println(logStr)
	global.Log.Println(logStr)
	router.Run(":" + port)
}
