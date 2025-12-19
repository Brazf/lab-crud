package org

import (
	"errors"
	repository "lab1-crud/internal/user"
	"lab1-crud/internal/user/model"
)

type OrgService interface {
	Create(org *model.Organization) error
	List() ([]model.Organization, error)
	Get(id uint) (*model.Organization, error)
	Update(id uint, data *model.Organization) error
	Delete(id uint) error
	GetUsersOrg(orgID uint) ([]model.OrganizationUser, error)
	AddUserToOrg(orgID, userID uint, role string) error
	RemoveUser(orgID, userID uint) error
}

type orgService struct {
	repo repository.OrgRepository
}

func NewOrgService(r repository.OrgRepository) OrgService {
	return &orgService{r}
}

func (s *orgService) Create(org *model.Organization) error {
	if org.Name == "" {
		return errors.New("nome da organização é obrigatório")
	}
	return s.repo.Create(org)
}

func (s *orgService) List() ([]model.Organization, error) {
	return s.repo.FindAll()
}

func (s *orgService) Get(id uint) (*model.Organization, error) {
	return s.repo.FindByID(id)
}

func (s *orgService) Update(id uint, data *model.Organization) error {
	org, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	org.Name = data.Name

	return s.repo.Update(org)
}

func (s *orgService) Delete(id uint) error {
	org, err := s.repo.FindByID(id)

	if err != nil {
		return err
	}

	if org != nil {
		return errors.New("Org não encontrada!")
	}

	return s.repo.Delete(id)
}

func (s *orgService) AddUserToOrg(orgID, userID uint, role string) error {
	orgUser := &model.OrganizationUser{
		OrganizationID: orgID,
		UserID:         userID,
		Role:           role,
	}
	return s.repo.AddUser(orgUser)
}

func (s *orgService) RemoveUser(orgID, userID uint) error {
	org, err := s.repo.FindByID(orgID)
	if err != nil {
		return err
	}

	if org != nil {
		return errors.New("Org não encontrada!")
	}

	userId, err := s.repo.FindUsersByOrg(userID)

	if err != nil {
		return err
	}

	if userId == nil {
		return errors.New("Usuário não encontrado!")
	}

	return s.repo.RemoveUser(orgID, userID)
}

func (s *orgService) GetUsersOrg(orgID uint) ([]model.OrganizationUser, error) {
	return s.repo.FindUsersByOrg(orgID)
}
