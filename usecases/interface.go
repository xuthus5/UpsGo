package usecases

import (
	"github.com/gin-gonic/gin"
)

type CoreInterface interface {
	//上传文件
	Upload(*gin.Context) (int, *Response)
	//删除文件
	Delete(*gin.Context) (int, *Response)
	//查询文件
	Get(*gin.Context) (int, *Response)
	//查询目录
	List(*gin.Context) (int, *List)
	//创建目录
	MakeDir(*gin.Context) (int, *Response)
}

type UserInterface interface {
	SignIn(*gin.Context) (int, *Response)
}
