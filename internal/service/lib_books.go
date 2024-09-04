package service

import (
	"library/internal/repository"
	"library/models"
)

type BookService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) Books {
	return &BookService{repo: repo}
}

func (s *BookService) GetAll() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) Create(book models.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) GetByID(id int) (models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) Update(book models.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) RentBook(userID, bookID int) error {
	return s.repo.RentBook(userID, bookID)
}

func (s *BookService) ReturnBook(userID, bookID int) error {
	return s.repo.ReturnBook(userID, bookID)
}
