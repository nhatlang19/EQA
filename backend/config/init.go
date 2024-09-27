package config

import (
	"EQA/backend/app/controller"
	"EQA/backend/app/pkg/mail"
	"EQA/backend/app/repository"
	"EQA/backend/app/service"

	"github.com/robfig/cron/v3"
)

type Initialization struct {
	userRepo repository.UserRepository
	authSvc  service.AuthService
	AuthCtrl controller.AuthController

	providerRepo repository.ProviderRepository

	mailSvc mail.MailService

	programRepo repository.ProgramRepository
	programSvc  service.ProgramService
	ProgramCtrl controller.ProgramController

	Cron *cron.Cron
}

func NewInitialization(
	userRepo repository.UserRepository,
	authService service.AuthService,
	authCtrl controller.AuthController,
	providerRepo repository.ProviderRepository,
	programRepo repository.ProgramRepository,
	programSvc service.ProgramService,
	programCtrl controller.ProgramController,
	mailSvc mail.MailService,
	cron *cron.Cron,
) *Initialization {
	return &Initialization{
		userRepo:     userRepo,
		authSvc:      authService,
		AuthCtrl:     authCtrl,
		providerRepo: providerRepo,
		programRepo:  programRepo,
		ProgramCtrl:  programCtrl,
		programSvc:   programSvc,
		mailSvc:      mailSvc,
		Cron:         cron,
	}
}
