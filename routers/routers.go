package routers

import (
	"github.com/gin-gonic/gin"
	"manger/routers/auth"
)

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
}
