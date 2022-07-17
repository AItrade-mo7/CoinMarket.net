package public

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("", Index)

	router.GET("/demo/get", Demo_get)
	router.POST("/demo/post", Demo_post)
}
