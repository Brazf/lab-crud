package user

import (
	"lab1-crud/internal/user/model"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
}

// ... (mantenha o UserRepository existente)

type OrgRepository interface {
	Create(org *model.Organization) error
	FindAll() ([]model.Organization, error)
	FindByID(id uint) (*model.Organization, error)
	Update(org *model.Organization) error
	Delete(id uint) error

	// Métodos para gerenciar usuários na org
	AddUser(orgUser *model.OrganizationUser) error
	RemoveUser(orgID, userID uint) error
	FindUsersByOrg(orgID uint) ([]model.OrganizationUser, error)
}
