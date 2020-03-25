package usecases

import "github.com/gin-gonic/gin"

type User struct {
}

func NewUser() UserInterface {
	return &User{}
}

func (this *User) SignIn(ctx *gin.Context) (int, *Response) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		//参数缺失
		return StatusClientError, &Response{
			Code:    ErrorParameterDefect,
			Message: "error",
		}
	}
	if username != conf.User.Username || password != conf.User.Password {
		//参数不匹配
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "error",
		}
	}
	//成功
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: conf.User.Token,
	}
}
