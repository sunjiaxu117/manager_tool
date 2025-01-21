package main

import (
	"github.com/gin-gonic/gin"
	"manger/db"
	"manger/routers"
)

func main() {
	//数据库
	db.InitDB()
	//服务器
	server := gin.Default()
	routers.RegisterRouters(server)
	server.Run(":8888")

}
