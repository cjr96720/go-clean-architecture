package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-clean-architecture/api/router"
	_ "go-clean-architecture/docs"
	"go-clean-architecture/internal/config"
)

//	@title			CRUD
//	@version		1.0
//	@description	API Endpoints

// @host	localhost:8080
// @BasePath
func main() {
	env := config.NewEnv()

	// Database
	db := config.NewPostgresConnection(env)

	// gin Engine
	gin := gin.Default()

	// add CORS
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: false,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
	}))

	// add swagger
	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Setup(db, gin)
	gin.Run(env.APIAddress)
}
