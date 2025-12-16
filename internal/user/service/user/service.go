package user

import (
	"errors"
	repository "lab1-crud/internal/user"
	"lab1-crud/internal/user/model"
	"strings"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Update(id uint, data *model.User) error
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Create(user *model.User) error {

	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if user.Email == "" || user.Name == "" {
		return errors.New("Email e Nome são obrigatórios!")
	}
	return s.repo.Create(user)
}

func (s *userService) Update(id uint, data *model.User) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.Name = data.Name
	user.Email = data.Email

	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("Usuário não encontrado!")
	}

	return s.repo.Delete(id)
}
