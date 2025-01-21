package auth

import (
	"github.com/gin-gonic/gin"
	"manger/controllers/auth"
)

func RegisterSubRouter(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	login(authGroup)
	loginout(authGroup)
}

func login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}

func loginout(authGroup *gin.RouterGroup) {
	authGroup.POST("/loginout", auth.Loginout)
}
