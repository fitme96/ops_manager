package routers

import (
	"ops_manager/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	router := r.Group("pveauth")
	router.POST("/", controller.Auth)

}
