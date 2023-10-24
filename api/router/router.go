package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-clean-architecture/api/controller"
	"go-clean-architecture/internal/repository"
	"go-clean-architecture/internal/service"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	baseRouter := gin.Group("")

	// healthCheck router
	healthCheckController := controller.NewHealthCheckController()
	baseRouter.GET("/healthz", healthCheckController.HealthCheck)

	v1apiRouter := gin.Group("/api/v1")

	// procuct router
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	v1apiRouter.POST("/product", productController.Add)
	v1apiRouter.GET("/product", productController.GetAll)
	v1apiRouter.GET("/product/:productId", productController.GetById)
	v1apiRouter.PATCH("/product/:productId", productController.Update)
	v1apiRouter.DELETE("/product/:productId", productController.Delete)
}
