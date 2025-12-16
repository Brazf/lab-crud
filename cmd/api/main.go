package main

import (
	router "lab1-crud/internal/http/router"
	"lab1-crud/internal/user/common/config"
	"lab1-crud/internal/user/common/database"
	handler "lab1-crud/internal/user/handler/user"
	"lab1-crud/internal/user/model"
	service "lab1-crud/internal/user/service/user"
	repository "lab1-crud/internal/user/storage/mysql/user"
	"log"
)

func main() {

	cfg := config.LoadConfig()

	database.ConnectDB(cfg)

	database.DB.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRoutes(userHandler)

	log.Println("Servidor está funcionando na porta:", cfg.AppPort)

	r.Run(":" + cfg.AppPort)

}
