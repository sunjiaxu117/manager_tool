// db/db.go
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type User struct {
	ID       uint `gorm:"primarykey;AUTO_INCREMENT"`
	Name     string
	Password string
}

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/manager?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
}
