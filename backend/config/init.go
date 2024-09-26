package config

import (
	"EQA/backend/app/controller"
	"EQA/backend/app/repository"
	"EQA/backend/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	authSvc  service.AuthService
	AuthCtrl controller.AuthController

	programRepo repository.ProgramRepository
	ProgramCtrl controller.ProgramController
}

func NewInitialization(
	userRepo repository.UserRepository,
	authService service.AuthService,
	authCtrl controller.AuthController,
	programRepo repository.ProgramRepository,
	programCtrl controller.ProgramController,
) *Initialization {
	return &Initialization{
		userRepo:    userRepo,
		authSvc:     authService,
		AuthCtrl:    authCtrl,
		programRepo: programRepo,
		ProgramCtrl: programCtrl,
	}
}
