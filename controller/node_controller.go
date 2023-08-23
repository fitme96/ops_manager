package controller

import (
	"ops_manager/service"

	"github.com/gin-gonic/gin"
)

func PveVersion(c *gin.Context) {
	if v, err := service.PveVersion(c); err != nil {
		c.JSON(200, "获取失败")
	} else {
		c.JSON(200, v)
	}
}

func Nodelist(c *gin.Context) {
	if nodelist, err := service.Nodelist(c); err != nil {
		c.JSON(200, "获取失败")
	} else {
		c.JSON(200, nodelist)
	}

}
func VmDetail(c *gin.Context) {
	var vm service.Vmcyclelife
	c.ShouldBindJSON(&vm)
	if vmdetail, err := service.VmDetail(c, vm); err != nil {
		c.JSON(200, "获取失败")
	} else {
		c.JSON(200, vmdetail)
	}
}

func VmList(c *gin.Context) {
	nodename := c.Param("nodename")
	if vmlist, err := service.VmList(c, nodename); err != nil {
		c.JSON(200, "获取失败")
	} else {
		c.JSON(200, vmlist)
	}
}
