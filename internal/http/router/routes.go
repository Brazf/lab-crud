package router

import (
	"lab1-crud/internal/http/middleware"
	orgHandler "lab1-crud/internal/user/handler/org"
	userHandler "lab1-crud/internal/user/handler/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *userHandler.UserHandler, orgHandler *orgHandler.OrgHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.GET("/org", orgHandler.ListOrgs)
	r.POST("/org", orgHandler.CreateOrg)

	orgGroup := r.Group("/org/:orgId")
	{
		orgGroup.GET("", middleware.Authorize("READ"), orgHandler.GetOrg)

		orgGroup.PUT("", middleware.Authorize("WRITE"), orgHandler.UpdateOrg)

		orgGroup.DELETE("", middleware.Authorize("ROOT"), orgHandler.DeleteOrg)

		orgGroup.POST("/users", middleware.Authorize("ROOT"), orgHandler.AddUserToOrg)

		orgGroup.DELETE("/users/:id", middleware.Authorize("ROOT"), orgHandler.RemoverUserOrg)

		orgGroup.GET("/users", middleware.Authorize("ROOT"), orgHandler.GetAllUsers)

	}

	return r
}
