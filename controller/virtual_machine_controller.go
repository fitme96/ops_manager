package controller

import (
	"ops_manager/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VmConfig(c *gin.Context) {
	var vmconfig service.Configvm
	var err error
	if err = c.ShouldBindJSON(&vmconfig); err != nil {
		c.JSON(200, "参数绑定失败")
	}
	if err = service.VmConfig(c, vmconfig); err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, "配置成功")

}

func ResizeDisk(c *gin.Context) {
	var resize service.Resize
	var err error
	if err = c.ShouldBindJSON(&resize); err != nil {
		c.JSON(200, "参数绑定失败")

	}
	if err = service.ResizeDisk(c, resize); err != nil {
		c.JSON(200, err)

	}
	c.JSON(200, "扩容成功")
}

func DeleteVm(c *gin.Context) {
	var vm service.Vmcyclelife
	c.ShouldBindJSON(&vm)
	if err := service.DeleteVm(c, vm); err != nil {
		c.JSON(200, err)

	}
	c.JSON(200, "删除成功")
}

func CloneVm(c *gin.Context) {
	var vmclone service.Clonevm
	c.ShouldBindJSON(&vmclone)
	newvm, err := service.CloneVm(c, vmclone)
	if err != nil {
		c.String(200, err.Error())

	}
	c.String(200, strconv.Itoa(newvm))

}

func VmStop(c *gin.Context) {
	var vm service.Vmcyclelife
	c.ShouldBindJSON(&vm)
	err := service.VmStop(c, vm)
	if err != nil {
		c.String(200, err.Error())

	}
	c.String(200, "停止成功")

}

func VmStart(c *gin.Context) {
	var vm service.Vmcyclelife
	c.ShouldBindJSON(&vm)
	err := service.VmStart(c, vm)
	if err != nil {
		c.String(200, err.Error())

	}
	c.String(200, "启动成功")
}

func VmShutdown(c *gin.Context) {
	var vm service.Vmcyclelife
	c.ShouldBindJSON(&vm)
	err := service.VmShutdown(c, vm)
	if err != nil {
		c.String(200, err.Error())

	}
	c.String(200, "关机成功")
}

func VmReboot(c *gin.Context) {
	var vm service.Vmcyclelife
	c.ShouldBindJSON(&vm)
	if err := service.VmReboot(c, vm); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, "重启完成")

	}
}

func NewSnapshot(c *gin.Context) {
	var snapshot service.Snapshotvm
	c.ShouldBindJSON(&snapshot)
	if err := service.NewSnapshot(c, snapshot); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, "快照创建完成")
	}
}

func Snapshots(c *gin.Context) {
	var snapshots service.Vmcyclelife
	c.ShouldBindJSON(&snapshots)
	if s, err := service.Snapshots(c, snapshots); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, s)
	}

}

func SnapshotRollback(c *gin.Context) {
	var snapshot service.Snapshotvm
	c.ShouldBindJSON(&snapshot)
	if err := service.SnapshotRollback(c, snapshot); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, "快照回滚完成")
	}
}
