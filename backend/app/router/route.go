package router

import (
	"EQA/backend/config"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// router.StaticFS("/", http.Dir("../frontend/.next"))

	// router.Use(middleware.JWTMiddleware())
	apiGroup := router.Group("/api")

	SetupAuthRouter(apiGroup, init.AuthCtrl)
	SetupProgram(apiGroup, init.ProgramCtrl)

	return router
}
