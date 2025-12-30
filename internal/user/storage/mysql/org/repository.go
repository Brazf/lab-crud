package org

import (
	"lab1-crud/internal/user"
	"lab1-crud/internal/user/model"

	"gorm.io/gorm"
)

type orgRepository struct {
	db *gorm.DB
}

func NewOrgRepository(db *gorm.DB) user.OrgRepository {
	return &orgRepository{db}
}

func (r *orgRepository) Create(org *model.Organization) error {
	return r.db.Create(org).Error
}

func (r *orgRepository) FindAll() ([]model.Organization, error) {
	var orgs []model.Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r *orgRepository) FindByID(id uint) (*model.Organization, error) {
	var org model.Organization
	err := r.db.First(&org, id).Error
	return &org, err
}

func (r *orgRepository) Update(org *model.Organization) error {
	return r.db.Save(org).Error
}

func (r *orgRepository) Delete(id uint) error {
	return r.db.Delete(&model.Organization{}, id).Error
}

func (r *orgRepository) AddUser(orgUser *model.OrganizationUser) error {
	return r.db.Save(orgUser).Error
}

func (r *orgRepository) RemoveUser(orgID, userID uint) error {
	return r.db.Delete(&model.OrganizationUser{}, "organization_id = ? AND user_id = ?", orgID, userID).Error
}

func (r *orgRepository) FindUsersByOrg(orgID uint) ([]model.OrganizationUser, error) {
	var usersOrg []model.OrganizationUser
	err := r.db.Where("organization_id = ?", orgID).Find(&usersOrg).Error

	if err != nil {
		return nil, err
	}

	return usersOrg, nil
}

func (r *orgRepository) SearchUserByOrg(orgID, userID uint) (*model.User, error) {
	return nil, nil
}
