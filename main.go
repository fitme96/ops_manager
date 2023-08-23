package main

import (
	"ops_manager/routers"
	"ops_manager/service"
)

func main() {

	r := routers.InitRouters()
	service.Setup()
	service.DbInit()

	r.Run(":8080")

}
