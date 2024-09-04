package service

import (
	"library/internal/repository"
	"library/models"
)

type AuthorService struct {
	repo repository.Authors
}

func NewAuthorsService(repo repository.Authors) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAll() ([]models.Author, error) {
	return s.repo.GetAll()
}

func (s *AuthorService) Create(author models.Author) error {
	return s.repo.Create(author)
}

func (s *AuthorService) GetByID(id int) (models.Author, error) {
	return s.repo.GetByID(id)
}

func (s *AuthorService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *AuthorService) Update(author models.Author) error {
	return s.repo.Update(author)
}
