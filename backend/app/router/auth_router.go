package router

import (
	"EQA/backend/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(c *gin.RouterGroup, ctrl controller.AuthController) {
	auth := c.Group("/auth")
	auth.POST("/login", ctrl.Login)
}
