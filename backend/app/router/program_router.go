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
	router.PUT("/:ID", ctrl.Update)
	router.PUT("/:ID/program_codes/:CodeId", ctrl.UpdateCode)
	router.DELETE("/:ID/program_codes/:CodeId", ctrl.DeleteCode)
	router.POST("/:ID/program_codes", ctrl.CreateCode)

	router.POST("/:ID/program_codes/:CodeId/details", ctrl.UpsertDatail)
}
