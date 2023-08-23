package routers

import (
	"ops_manager/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	router := r.Group("user")
	router.POST("/create", controller.CreateUser)

}
