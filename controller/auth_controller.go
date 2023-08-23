package controller

import (
	"ops_manager/service"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var user service.User
	c.ShouldBind(&user)
	if token, err := service.Auth(c, user); err != nil {
		c.JSON(200, service.Response{Code: 401, Data: nil, Msg: "鉴权失败"})
	} else {
		c.JSON(200, service.Response{Code: 200, Data: token, Msg: "鉴权成功"})
	}

}
