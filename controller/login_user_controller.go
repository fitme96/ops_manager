package controller

import (
	"ops_manager/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user service.User
	c.ShouldBindJSON(&user)
	if username, b := service.CreateUser(c, user); b {
		c.JSON(200, gin.H{"msg": username + "创建成功"})

	} else {
		c.JSON(200, gin.H{"msg": username + "用户已存在"})

	}
}
