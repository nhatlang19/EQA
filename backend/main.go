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

// @title           BLUE HOUSE API
// @version         1.0
// @description     This page allows you to explore and interact with BLUE HOUSE System REST API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  nguyennb@teecom.vn
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @host      localhost:8080
// @BasePath  /

func main() {
	port := os.Getenv("PORT")
	docs.SwaggerInfo.BasePath = "/api/"

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
