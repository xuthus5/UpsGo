package delivers

import (
	"UpsGo/usecases"
	"github.com/gin-gonic/gin"
)

var user = usecases.NewUser()

func SignIn(this *gin.Context) {
	code, res := user.SignIn(this)
	this.JSON(code, res)
}
