package middleware

import (
	"UpsGo/utils"
	"github.com/gin-gonic/gin"
)

const (
	//token缺失
	TokenDefect = 4002
	//token无效
	TokenTampered = 4003
)

func TokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("token")
		if auth == "" {
			context.JSON(400, gin.H{
				"code":    TokenDefect,
				"message": "Request header has no token field",
			})
			context.Abort()
			return
		}
		var conf = utils.GetConfig()
		if auth != conf.User.Token {
			context.JSON(400, gin.H{
				"code":    TokenTampered,
				"message": "Token matching failed",
			})
			context.Abort()
			return
		}
	}
}
