package usecases

import (
	"UpsGo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/upyun/go-sdk/upyun"
)

var conf = utils.GetConfig()

// Core 操作服务核心
type Core struct {
}

func NewCore() CoreInterface {
	return &Core{}
}

func (this *Core) Upload(ctx *gin.Context) (int, *Response) {
	header, err := ctx.FormFile("file")
	path := ctx.PostForm("path")
	if err != nil || path == "" || header == nil {
		//上传文件为空
		return StatusClientError, &Response{
			Code:    ErrorParameterDefect,
			Message: "ErrorParameterDefect",
		}
	}
	dst := header.Filename
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   conf.Ups.Bucket,
		Operator: conf.Ups.Operator,
		Password: conf.Ups.Password,
	})

	source, _ := header.Open()
	if err := up.Put(&upyun.PutObjectConfig{
		Path:   path + "/" + dst,
		Reader: source,
	}); err != nil {
		//上传失败
		return StatusServerError, &Response{
			Code:    ErrorUpload,
			Message: "ErrorUpload:" + err.Error(),
		}
	}
	//成功
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this *Core) Delete(ctx *gin.Context) (int, *Response) {
	path := ctx.Query("path")

	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   conf.Ups.Bucket,
		Operator: conf.Ups.Operator,
		Password: conf.Ups.Password,
	})

	if err := up.Delete(&upyun.DeleteObjectConfig{
		Path:  path,
		Async: false,
	}); err != nil {
		//删除失败
		return StatusServerError, &Response{
			Code:    ErrorDelete,
			Message: "ErrorDelete:" + err.Error(),
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this *Core) Get(ctx *gin.Context) (int, *Response) {
	panic("get")
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this *Core) List(ctx *gin.Context) (int, *List) {
	path := ctx.Query("path")
	if path == "" {
		path = "/"
	}
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   conf.Ups.Bucket,
		Operator: conf.Ups.Operator,
		Password: conf.Ups.Password,
	})
	objsChan := make(chan *upyun.FileInfo, 10)

	go func() {
		err := up.List(&upyun.GetObjectsConfig{
			Path:        path,
			ObjectsChan: objsChan,
		})
		if err != nil {
			fmt.Println(err)
		}
	}()

	var list []*upyun.FileInfo

	for obj := range objsChan {
		list = append(list, obj)
	}

	return StatusOK, &List{
		Code:    StatusOK,
		Message: conf.Ups.Domain,
		Data:    list,
		Count:   len(list),
	}
}

func (this *Core) MakeDir(ctx *gin.Context) (int, *Response) {
	dir := ctx.Query("dir")
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   conf.Ups.Bucket,
		Operator: conf.Ups.Operator,
		Password: conf.Ups.Password,
	})

	if err := up.Mkdir(dir); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorReadRemote,
			Message: err,
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}
