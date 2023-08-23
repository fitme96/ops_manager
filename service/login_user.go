package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, user User) (string, bool) {
	var exsit User
	a := db.First(&exsit, "username =?", user.Username)
	if errors.Is(a.Error, gorm.ErrRecordNotFound) {
		hashdPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		newUser := User{Username: user.Username, Password: string(hashdPassword)}
		db.Create(&newUser)
		return user.Username, true

	}

	return user.Username, false
}
