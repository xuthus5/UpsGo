package delivers

import (
	"UpsGo/usecases"
	"github.com/gin-gonic/gin"
)

var core = usecases.NewCore()

// List 文件列表
func List(this *gin.Context) {
	this.JSON(core.List(this))
}

// Upload 上传文件
func Upload(this *gin.Context) {
	this.JSON(core.Upload(this))
}

// MakeDir 创建目录
func MakeDir(this *gin.Context) {
	this.JSON(core.MakeDir(this))
}

// Delete 删除文件、目录
func Delete(this *gin.Context) {
	this.JSON(core.Delete(this))
}
