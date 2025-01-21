package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"manger/db"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user UserInfo
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var users db.User
	users.Name = user.Username
	users.Password = user.Password
	result := db.DB.First(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(401, gin.H{"message": "invalid credentials"})
	} else if result.Error != nil {
		c.JSON(500, gin.H{"message": "server error"})
	} else {
		c.JSON(200, gin.H{"message": "login success"})
	}
}

func Loginout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
