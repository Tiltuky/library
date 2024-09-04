package service

import (
	"library/internal/repository"
	"library/models"
)

type UserService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) Users {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) Create(user models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *UserService) Update(user models.User) error {
	return s.repo.Update(user)
}
