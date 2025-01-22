// db/db.go
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

var (
	userName  string = "root"
	password  string = "root"
	ipAddress string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "manager"
	charset   string = "utf8mb4"
)

func InitDB() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	//此处使用musql驱动 则需要在import中加入mysql的初始化驱动 并不使用但是需要进行mysql的初始化
	//_ "github.com/go-sql-driver/mysql"
	var err error
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
}
