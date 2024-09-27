package router

import (
	"EQA/backend/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupProgram(c *gin.RouterGroup, ctrl controller.ProgramController) {
	router := c.Group("/programs")
	router.GET("", ctrl.GetAll)
	router.GET("/:ID", ctrl.GetOne)
	router.GET("/reminder", ctrl.Reminder)
}
