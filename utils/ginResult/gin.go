package ginResult

import "github.com/EasyGolang/goTools/mGin"

/*
# 使用方法
engine.GET("/data", func(c *gin.Context) {
	res := struct {
		Name  string `bson:"name"`
		Age   int    `bson:"age"`
		Email string `bson:"email"`
	}{
		Name:  "小明",
		Age:   18,
		Email: "110@qq.com",
	}
	c.JSON(http.StatusOK, result.OK.WithData(res))
})

engine.GET("/common/err", func(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, result.Err.WithMsg("通用错误"))
})
*/

var (
	// 通用
	OK   = mGin.Response(200, "Succeed") // 通用成功
	Fail = mGin.Response(500, "Fail")    // 通用错误
	Hz   = mGin.Response(208, "请求太频繁")

	// 模块级错误码 - 用户模块
	DBErr           = mGin.Response(199, "数据库出错")
	LoginSucceed    = mGin.Response(200, "登录成功")
	RegisterSucceed = mGin.Response(200, "注册成功")
	CodeSucceed     = mGin.Response(200, "验证码已发送")

	ErrEmail     = mGin.Response(201, "邮箱格式不正确")
	UserExist    = mGin.Response(202, "用户已存在")
	CodeFail     = mGin.Response(203, "验证码发送失败，请稍后再试")
	CodeErr      = mGin.Response(204, "验证码错误")
	RegisterFail = mGin.Response(205, "注册失败,请重新尝试")
	LoginFail    = mGin.Response(206, "登录失败,请重新尝试")
	UserAuthErr  = mGin.Response(207, "身份验证失败")

	AuthErr = mGin.Response(260, "接口授权错误")
)
