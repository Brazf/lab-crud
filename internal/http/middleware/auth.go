package middleware

import (
	"lab1-crud/internal/user/common/database"
	"lab1-crud/internal/user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Authorize(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		orgIDStr := c.Param("orgId")
		userIDStr := c.GetHeader("X-User-ID")

		if userIDStr == "" || orgIDStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UserID ou orgID não recebido"})
			c.Abort()
			return
		}

		uID, _ := strconv.Atoi(userIDStr)
		oID, _ := strconv.Atoi(orgIDStr)

		var orgUser model.OrganizationUser
		err := database.DB.Where("organization_id = ? AND user_id = ?", oID, uID).First(&orgUser).Error

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Usuário não encontrado na organização!"})
		}

		authorized := false

		for _, role := range requiredRoles {
			if orgUser.Role == role {
				authorized = true
				break
			}
		}

		if orgUser.Role == model.RoleRoot {
			authorized = true
		} else if orgUser.Role == model.RoleWrite {
			for _, r := range requiredRoles {
				if r == model.RoleRead {
					authorized = true
				}
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "usuário sem autorizações"})
			c.Abort()
			return
		}

		c.Next()

	}
}
