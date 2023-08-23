package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	apigroup := r.Group("api")

	VmRouter(apigroup)
	AuthRouter(apigroup)
	UserRouter(apigroup)
	return r

}
