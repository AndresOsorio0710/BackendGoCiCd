package main

import (
	"log"

	_ "github.com/AndresOsorio0710/BackendGoCiCd/docs"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/application/services"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/config"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/dbcontext"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/infrastructure/repository"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/presentation/handlers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task API
// @version 1.0
// @description Esta API permite manejar tareas.
// @host localhost:8080
// @BasePath /api
// @schemes http

func main() {
	err := config.LoadConfig("../internals/config")
	if err != nil {
		log.Fatal("No se pudo cargar la configuración:", err)
	}

	dbCtx, err := dbcontext.NewDbContext(config.Cfg.Postgres)
	if err != nil {
		log.Fatal("Error al inicializar la base de datos:", err)
	}
	taskRepo := repository.NewTaskRepository(dbCtx)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := gin.Default()

	taskHandler.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
