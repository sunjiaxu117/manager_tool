package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"manger/log"
	"time"

	"manger/db"
	"manger/routers"
)

func main() {
	log.InitLogger()
	db.InitDB()
	//服务器
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"}, // 允许访问的前端域名，包括端口
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routers.RegisterRouters(server)
	server.Run(":8888")

}
