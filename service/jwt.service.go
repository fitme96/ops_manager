package service

import (
	"ops_manager/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Auth(c *gin.Context, user User) (string, error) {
	var dbuser User
	db.First(&dbuser, "username = ?", user.Username)
	if err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	//生成token
	tokenString, _ := middleware.GenToken(user.Username)

	return tokenString, nil

}
