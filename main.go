package main

import (
	"github.com/gin-gonic/gin"
	"manger/log"

	"manger/db"
	"manger/routers"
)

func main() {
	log.InitLogger()
	db.InitDB()
	//服务器
	server := gin.Default()
	routers.RegisterRouters(server)
	server.Run(":8888")

}
