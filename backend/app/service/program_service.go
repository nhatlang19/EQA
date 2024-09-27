package service

import (
	dto "EQA/backend/app/domain/dto"
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
)

type ProgramService interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Reminder()
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
	id, _ := strconv.Atoi(c.Param("ID"))
	data, err := s.programRepo.FindOne(id)
	if err != nil {
		pkg.PanicException(constant.DataNotFound, fmt.Sprintf("Not found program with ID = %d", id))
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
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
