package service

import (
	dto "EQA/backend/app/domain/dto"
	"EQA/backend/app/domain/model"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"EQA/backend/app/constant"
	"EQA/backend/app/pkg"
	"EQA/backend/app/pkg/mail"
	"EQA/backend/app/repository"

	exporter "EQA/backend/app/pkg/exporter"
)

type ProgramService interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Reminder()
	Update(c *gin.Context)
	UpdateCode(c *gin.Context)
	CreateCode(c *gin.Context)
	DeleteCode(c *gin.Context)
	UpsertDatail(c *gin.Context)
	Export(c *gin.Context)
}

type ProgramServiceImpl struct {
	programRepo repository.ProgramRepository
	mailSvc     mail.MailService
}

func ProgramServiceInit(programRepo repository.ProgramRepository, mailSvc mail.MailService) *ProgramServiceImpl {
	return &ProgramServiceImpl{
		programRepo: programRepo,
		mailSvc:     mailSvc,
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

func (s ProgramServiceImpl) GetOne(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var filter dto.ProgramCodeFilter
	filter.Year, _ = strconv.Atoi(c.DefaultQuery("year", "-1"))

	id, _ := strconv.Atoi(c.Param("ID"))
	data, err := s.programRepo.FindOneWithFilter(id, filter)
	if err != nil {
		pkg.PanicException(constant.DataNotFound, fmt.Sprintf("Not found program with ID = %d", id))
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s ProgramServiceImpl) Update(c *gin.Context) {
	defer pkg.PanicHandler(c)
	id, _ := strconv.Atoi(c.Param("ID"))
	var req dto.UpdateProgramResp
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	res, err := s.programRepo.FindOne(id)
	if err != nil {
		pkg.PanicException(constant.DataNotFound, "Invalid Id")
	}
	res.Name = req.Name
	res, err = s.programRepo.Save(&res)
	if err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (s ProgramServiceImpl) UpdateCode(c *gin.Context) {
	defer pkg.PanicHandler(c)
	id, _ := strconv.Atoi(c.Param("ID"))
	codeId, _ := strconv.Atoi(c.Param("CodeId"))
	var req dto.UpdateProgramCodeResp
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	res, err := s.programRepo.FindByCodeId(id, codeId)
	if err != nil {
		pkg.PanicException(constant.DataNotFound, "Invalid Id")
	}
	res.Name = req.Name
	res.Year = req.Year
	res, err = s.programRepo.SaveCode(&res)
	if err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (s ProgramServiceImpl) CreateCode(c *gin.Context) {
	defer pkg.PanicHandler(c)
	id, _ := strconv.Atoi(c.Param("ID"))
	var req dto.CreateProgramCodeResp
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	res, err := s.programRepo.CreateCode(&model.ProgramCode{Name: req.Name, ProgramId: id, Year: req.Year})
	if err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, res))
}

func (s ProgramServiceImpl) DeleteCode(c *gin.Context) {
	defer pkg.PanicHandler(c)
	codeId, _ := strconv.Atoi(c.Param("CodeId"))
	var err error
	err = s.programRepo.DeleteCodeById(codeId)
	if err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.BadRequest, err.Error())
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Successfully deleted"))
}

func (s ProgramServiceImpl) UpsertDatail(c *gin.Context) {
	defer pkg.PanicHandler(c)
	codeId, _ := strconv.Atoi(c.Param("CodeId"))
	var req dto.ProgramCodeReminderReq
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		pkg.PanicException(constant.BadRequest, err.Error())
	}
	err = s.programRepo.DeleteDetailCodeById(codeId)
	if err != nil {
		fmt.Println(err)
		pkg.PanicException(constant.BadRequest, err.Error())
	}
	for _, data := range req.ProgramCodeReminders {
		_, err := s.programRepo.SaveCodeDetail(&model.ProgramCodeReminder{ProgramCodeId: codeId, Sample: data.Sample, DateOfReceive: data.DateOfReceive, DateOfReturn: data.DateOfReturn, Status: data.Status, PercentPassed: data.PercentPassed})
		if err != nil {
			fmt.Println(err)
			pkg.PanicException(constant.BadRequest, err.Error())
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Success"))
}

func (s ProgramServiceImpl) Reminder() {
	fmt.Println("Cron job running every minute:", time.Now())
	programs, err := s.programRepo.FindForReminder(3)
	if err != nil {
		fmt.Println("Failed to fetch data from database")
	}

	fmt.Println("Sending mail .....:")
	for _, program := range programs {
		payload := dto.EmaiReminderPayload{
			Name:     program.ProgramName,
			Code:     program.Code,
			Deadline: program.DateOfReturn.Format("01/02/2006"),
			Sample:   program.Sample,
		}

		subject := "Nhắc nhở thực hiện ngoại kiểm"
		templatePath := "./backend/mail_templates/reminder_near_date_of_return.html"
		err = s.mailSvc.SendEmail(os.Getenv("MAIL_TO"), subject, os.Getenv("MAIL_BCC"), templatePath, payload)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (s ProgramServiceImpl) Export(c *gin.Context) {
	defer pkg.PanicHandler(c)
	data, err := s.programRepo.FindForExport()
	if err != nil {
		pkg.PanicException(constant.InternalServer, "Failed to fetch data from database")
	}

	exporter, err := exporter.ExporterFactory(exporter.ProgramType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	buffer, err := exporter.Export(data)
	if err != nil {
		fmt.Println("Error exporting:", err)
		return
	}

	downloadName := time.Now().UTC().Format("export-data-program.xlsx")

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())
}
