package controller

import (
	"EQA/backend/app/service"

	"github.com/gin-gonic/gin"
)

type ProgramController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Reminder(c *gin.Context)
	Update(c *gin.Context)
	UpdateCode(c *gin.Context)
	CreateCode(c *gin.Context)
	UpsertDatail(c *gin.Context)
	DeleteCode(c *gin.Context)
	Export(c *gin.Context)
}

type ProgramControllerImpl struct {
	svc service.ProgramService
}

func ProgramControllerInit(service service.ProgramService) *ProgramControllerImpl {
	return &ProgramControllerImpl{
		svc: service,
	}
}

// Program PingExample
// @Summary Get All
// @Schemes
// @Description Get Program Data
// @Tags Program
// @Accept json
// @Produce json
// @securityDefinitions.apiKey Authorization
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {object} dto.ApiResponse
// @Router /programs [get]
// @Param source_id query string false "source_id"
// @Param rating_id query string false "rating_id"
// @Param month query string false "month"
// @Param year query string false "year"
// @Param page query string false "Page description"
// @Param limit query string false "Limit description"
// @Param search query string false "Search description"
func (bh ProgramControllerImpl) GetAll(c *gin.Context) {
	bh.svc.GetAll(c)
}

// @Summary Get program
// @Schemes
// @Description Get program by ID
// @Tags Program
// @Accept json
// @Produce json
// @securityDefinitions.apiKey Authorization
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {object} model.Program
// @Router /programs/{ID} [get]
func (bh ProgramControllerImpl) GetOne(c *gin.Context) {
	bh.svc.GetOne(c)
}

func (bh ProgramControllerImpl) Reminder(c *gin.Context) {
	bh.svc.Reminder()
}

// Put PingExample program
// @Summary Get update program
// @Schemes
// @Description Put program Data
// @Tags Program
// @Accept json
// @Produce json
// @securityDefinitions.apiKey Authorization
// @in header
// @name Authorization
// @Security JWT
// @Param data body model.Program
// @Success 200 {object} dto.ApiResponse
// @Router /programs/{ID} [put]

func (bh ProgramControllerImpl) Update(c *gin.Context) {
	bh.svc.Update(c)
}

func (bh ProgramControllerImpl) UpdateCode(c *gin.Context) {
	bh.svc.UpdateCode(c)
}

func (bh ProgramControllerImpl) CreateCode(c *gin.Context) {
	bh.svc.CreateCode(c)
}

func (bh ProgramControllerImpl) UpsertDatail(c *gin.Context) {
	bh.svc.UpsertDatail(c)
}

func (bh ProgramControllerImpl) DeleteCode(c *gin.Context) {
	bh.svc.DeleteCode(c)
}

func (bh ProgramControllerImpl) Export(c *gin.Context) {
	bh.svc.Export(c)
}
