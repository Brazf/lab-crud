package main

import (
	"lab1-crud/internal/config"
	"lab1-crud/internal/database"
	"lab1-crud/internal/handler"
	"lab1-crud/internal/model"
	"lab1-crud/internal/repository"
	"lab1-crud/internal/service"
	"log"
)

func main() {

	cfg := config.LoadConfig()

	database.ConnectDB(cfg)

	database.DB.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := handler.SetupRoutes(userHandler)

	log.Println("Servidor está funcionando na porta:", cfg.AppPort)

	r.Run(":" + cfg.AppPort)

}
