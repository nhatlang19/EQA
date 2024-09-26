package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"EQA/backend/app/constant"
	"EQA/backend/app/pkg"
	"EQA/backend/app/repository"
)

type ProgramService interface {
	GetAll(c *gin.Context)
}

type ProgramServiceImpl struct {
	programRepo repository.ProgramRepository
}

func ProgramServiceInit(programRepo repository.ProgramRepository) *ProgramServiceImpl {
	return &ProgramServiceImpl{
		programRepo: programRepo,
	}
}

func (s ProgramServiceImpl) GetAll(c *gin.Context) {
	defer pkg.PanicHandler(c)
	data, err := s.programRepo.FindAll()
	if err != nil {
		pkg.PanicException(constant.InternalServer, "Failed to fetch data from database")
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}
