package main

import (
	router "lab1-crud/internal/http/router"
	"lab1-crud/internal/user/common/config"
	"lab1-crud/internal/user/common/database"
	userHandler "lab1-crud/internal/user/handler/user"
	"lab1-crud/internal/user/model"
	userService "lab1-crud/internal/user/service/user"
	userRepository "lab1-crud/internal/user/storage/mysql/user"
	"log"

	orgHandler "lab1-crud/internal/user/handler/org"
	orgService "lab1-crud/internal/user/service/org"
	orgRepository "lab1-crud/internal/user/storage/mysql/org"
)

func main() {

	cfg := config.LoadConfig()

	database.ConnectDB(cfg)

	database.DB.AutoMigrate(&model.User{}, &model.Organization{}, &model.OrganizationUser{})

	userRepo := userRepository.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	orgRepo := orgRepository.NewOrgRepository(database.DB)
	orgService := orgService.NewOrgService(orgRepo)
	orgHandler := orgHandler.NewOrgHandler(orgService)

	r := router.SetupRoutes(userHandler, orgHandler)

	log.Println("Servidor está funcionando na porta:", cfg.AppPort)

	r.Run(":" + cfg.AppPort)

}
