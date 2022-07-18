package public

import (
	"net/http"
	"time"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// webSocket请求ping 返回pong
func WsServer(c *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// go func() {
	// 	for {
	// 		// 读取ws中的数据
	// 		mt, message, err := ws.ReadMessage()
	// 		if err != nil {
	// 			break
	// 		}
	// 		if mStr.ToStr(message) == "ping" {
	// 			message = []byte("pong")
	// 		}
	// 		// 写入ws数据
	// 		err = ws.WriteMessage(mt, message)
	// 		if err != nil {
	// 			break
	// 		}
	// 	}
	// }()
	go func() {
		for {
			data := map[string]any{
				"France": "Paris",
				"Italy":  "Rome",
				"Japan":  "Tokyo",
				"India":  mCount.GetRound(100, 99999),
			}

			b, _ := jsoniter.Marshal(data)
			ws.WriteMessage(1, b)
			time.Sleep(time.Second) // 一秒执行一次
		}
	}()
}
