package public

import (
	"net/http"
	"strconv"

	"CoinMarket.net/utils/ginResult"
	"github.com/gin-gonic/gin"
)

/*

param := c.Query("param") // 查询请求URL后面的参数
id := c.Query("id") //查询请求URL后面的参数
page := c.DefaultQuery("page", "0") //查询请求URL后面的参数，如果没有填写默认值
name := c.PostForm("name") //从表单中查询参数

//POST和PUT主体参数优先于URL查询字符串值。
name := c.Request.FormValue("name")

//返回POST并放置body参数，URL查询参数被忽略
name := c.Request.PostFormValue("name")

//从表单中查询参数，如果没有填写默认值
message := c.DefaultPostForm("message", "aa")


*/
// get
func Demo_get(c *gin.Context) {
	rawQuery := make(map[string]string)
	c.Bind(&rawQuery)

	n, _ := strconv.Atoi(c.Query("age"))
	json := struct {
		Name string `bson:"Name"`
		Age  int    `bson:"Age"`
		Data string `bson:"Data"`
	}{
		Name: c.Query("name"),
		Age:  n,
	}
	json.Data = "get 请求返回"
	json.Name = json.Name + "@GET, 这里是 CoinMarket.net"

	c.JSON(http.StatusOK, ginResult.OK.WithData(json))
}

// post
func Demo_post(c *gin.Context) {
	// 两个只能存在一个
	// rawQuery := make(map[string]interface{})
	// c.ShouldBind(&rawQuery)

	var json struct {
		Name string `bson:"Name"`
		Age  int    `bson:"Age"`
		Data string `bson:"Data"`
	}
	c.ShouldBind(&json)
	json.Data = "post 请求返回"
	json.Name = json.Name + "@POST,  这里是 CoinMarket.net"

	c.JSON(http.StatusOK, ginResult.OK.WithData(json))
}

// https://zhuanlan.zhihu.com/p/377795123
