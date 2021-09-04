package router

import (
	"github.com/gin-gonic/gin"
	UserControl "go_v/control"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", UserControl.Register)

	r.POST("/api/user/login", UserControl.Login)

	return r
}
