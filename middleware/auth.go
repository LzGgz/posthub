package middleware

import (
	"posthub/controller"
	"posthub/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		controller.ResponseError(c, controller.CodeNeedLogin)
		c.Abort()
		return
	}
	parts := strings.SplitN(auth, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		controller.ResponseError(c, controller.CodeTokenInvaildFormat)
		c.Abort()
		return
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	mc, err := jwt.ParseToken(parts[1])
	if err != nil {
		controller.ResponseError(c, controller.CodeInvalidToken)
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	c.Set(controller.CtxId, mc.ID)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
