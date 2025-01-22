package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manger/model"
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
	m := &model.Users{}
	err := m.GetByName(user.Username)
	if err != nil {
		fmt.Println(err)
	}
	if m.Name == user.Username && m.Password == user.Password {
		c.JSON(200, gin.H{"message": "login success"})
	} else {
		c.JSON(400, gin.H{"message": "login fail"})
	}
}

func Loginout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
