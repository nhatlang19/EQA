//go:build wireinject
// +build wireinject

package config

import (
	"EQA/backend/app/controller"
	"EQA/backend/app/repository"
	"EQA/backend/app/service"

	"EQA/backend/app/pkg/cloud/sqs"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var sqsClient = wire.NewSet(sqs.NewSQS,
	wire.Bind(new(sqs.MessageClient), new(*sqs.SQS)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var authServiceSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
)

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var programRepoSet = wire.NewSet(repository.ProgramRepositoryInit,
	wire.Bind(new(repository.ProgramRepository), new(*repository.ProgramRepositoryImpl)),
)

var programServiceSet = wire.NewSet(service.ProgramServiceInit,
	wire.Bind(new(service.ProgramService), new(*service.ProgramServiceImpl)),
)

var programCtrlSet = wire.NewSet(controller.ProgramControllerInit,
	wire.Bind(new(controller.ProgramController), new(*controller.ProgramControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization,
		db,
		userRepoSet,
		authServiceSet,
		authCtrlSet,
		programRepoSet,
		programServiceSet,
		programCtrlSet,
	)
	return nil
}
