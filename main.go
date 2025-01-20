package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"manger/utils"
	"net/http"
)

type User struct {
	ID       uint `gorm:"primarykey;AUTO_INCREMENT"`
	Name     string
	Password string
}

func main() {
	s, e := utils.GenToken("sun")
	fmt.Println(s, e)
	//数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/manager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	//
	db.AutoMigrate(&User{})
	//服务器
	server := gin.Default()

	//注册

	server.POST("/user/register", func(context *gin.Context) {
		u := User{}
		context.BindJSON(&u)
		//查询用户是否存在
		res := db.Where("name = ?", u.Name).First(&User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册失败,用户名存在",
			})
		} else {
			db.Create(&u)
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册成功",
			})
		}
	})

	server.Run(":8888")

}
