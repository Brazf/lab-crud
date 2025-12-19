package org

import (
	"lab1-crud/internal/user/model"
	service "lab1-crud/internal/user/service/org"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrgHandler struct {
	s service.OrgService
}

func NewOrgHandler(s service.OrgService) *OrgHandler {
	return &OrgHandler{s}
}

func (h *OrgHandler) CreateOrg(c *gin.Context) {
	var org model.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.Create(&org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, org)
}

func (h *OrgHandler) ListOrgs(c *gin.Context) {
	orgs, err := h.s.List()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orgs)

}

func (h *OrgHandler) GetOrg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("orgId"))

	org, err := h.s.Get(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, org)

}

func (h *OrgHandler) UpdateOrg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("orgID"))

	var org model.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"erro": err.Error()})
		return
	}

	if err := h.s.Update(uint(id), &org); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "org atualizado"})

}

func (h *OrgHandler) DeleteOrg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("orgId"))

	if err := h.s.Delete(uint(id)); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "org deletada"})
}

func (h *OrgHandler) AddUserToOrg(c *gin.Context) {
	orgID, _ := strconv.Atoi(c.Param("orgId"))
	var input struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.AddUserToOrg(uint(orgID), input.UserID, input.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "usuário adicionado com sucesso"})
}

func (h *OrgHandler) RemoverUserOrg(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	orgId, _ := strconv.Atoi(c.Param("orgId"))

	if err := h.s.RemoveUser(uint(orgId), uint(userId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "usuário removido com sucesso"})

}

func (h *OrgHandler) GetAllUsers(c *gin.Context) {
	idOrg, _ := strconv.Atoi(c.Param("orgId"))
	users, err := h.s.GetUsersOrg(uint(idOrg))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
