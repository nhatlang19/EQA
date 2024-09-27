package main

import (
	"EQA/backend/app/router"
	"EQA/backend/config"
	"os"

	docs "EQA/backend/docs"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

// @title           EQA API
// @version         1.0
// @description     This page allows you to explore and interact with EQA System REST API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  nhatlang19@gmail.com
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @host      localhost:8081
// @BasePath  /

func main() {
	port := os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = "/api/"

	init := config.Init()
	go init.Cron.Start()
	defer init.Cron.Stop()

	app := router.Init(init)

	app.Run("0.0.0.0:" + port)
}
