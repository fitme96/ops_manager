package service

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func DbInit() {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DatabaseSetting.User,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.Name,
		),
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

}
