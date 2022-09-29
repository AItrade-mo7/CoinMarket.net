package router

import (
	"os"
	"time"

	"CoinMarket.net/server/global"
	"CoinMarket.net/server/global/config"
	"CoinMarket.net/server/router/api"
	"CoinMarket.net/server/router/middle"
	"CoinMarket.net/server/router/private"
	"CoinMarket.net/server/router/public"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	// 加载日志文件
	fileName := config.Dir.Log + "/HTTP-T" + time.Now().Format("06年1月02日15时") + ".log"
	logFile, _ := os.Create(fileName)
	/*
		加载模板
		https://www.gouguoyin.cn/posts/10103.html
	*/

	// 创建服务
	app := fiber.New(fiber.Config{
		ServerHeader: "CoinMarket.net",
	})

	// 跨域
	app.Use(cors.New())
	// 限流
	app.Use(limiter.New(limiter.Config{
		Max:        200,
		Expiration: 1 * time.Second,
	}))
	// 日志中间件
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${status} - ${method} ${latency} ${path} \n",
		TimeFormat: "2006-01-02 - 15:04:05",
		Output:     logFile,
	}), middle.Public, compress.New(), favicon.New())

	// CoinMarket
	A_api := app.Group("/CoinMarket")

	// /api/public
	public.Router(A_api)

	// /api/private
	private.Router(A_api)

	// Ping
	app.Use(api.Ping)

	listenHost := mStr.Join(":", config.AppInfo.Port)
	global.Log.Println(mStr.Join(`启动服务: http://127.0.0.1`, listenHost))
	app.Listen(listenHost)
}
