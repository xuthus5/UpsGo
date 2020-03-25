package delivers

import (
	"UpsGo/middleware"
	"UpsGo/utils"
	"github.com/gin-gonic/gin"
)

// Run 程序启动的入口
func Run() {
	var conf = utils.GetConfig()
	var router = InitRouter()
	//测试路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//关闭debug
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	//服务运行
	_ = router.Run(conf.Server)
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	//开启允许跨域
	router.Use(middleware.CORS())
	//用户路由
	var user = router.Group("/api/user")
	{
		//登录
		user.POST("/signin", SignIn)
	}
	//通用路由
	var core = router.Group("/api/core", middleware.TokenAuth())
	{
		//列表
		core.GET("/list", List)
		//创建目录
		core.GET("/mkdir", MakeDir)
		//删除文件、目录
		core.GET("/delete", Delete)
		//上传文件
		core.POST("/upload", Upload)
	}
	return router
}
