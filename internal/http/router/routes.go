package router

import (
	handler "lab1-crud/internal/user/handler/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	return r
}
