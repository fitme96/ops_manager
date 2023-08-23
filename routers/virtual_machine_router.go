package routers

import (
	"ops_manager/controller"
	"ops_manager/middleware"

	"github.com/gin-gonic/gin"
)

func VmRouter(r *gin.RouterGroup) {

	router := r.Group("vm").Use(middleware.JWTAuthMiddleware())
	{
		router.GET("/version", controller.PveVersion)
		router.POST("/detail", controller.VmDetail)
		router.GET("/list/:nodename", controller.VmList)
		router.GET("/nodelist", controller.Nodelist)

		// 虚拟机配置
		router.POST("/clone", controller.CloneVm)
		router.POST("/resize", controller.ResizeDisk)
		router.POST("/config", controller.VmConfig)

		//生命周期管理
		router.POST("/delete", controller.DeleteVm)
		router.POST("/stop", controller.VmStop)
		router.POST("/start", controller.VmStart)
		router.POST("/shutdown", controller.VmShutdown)
		router.POST("/reboot", controller.VmReboot)
		// 快照管理
		router.POST("/newsnapshot", controller.NewSnapshot)
		router.POST("/snapshots", controller.Snapshots)
		router.POST("/snapshotrollback", controller.SnapshotRollback)

	}

}
