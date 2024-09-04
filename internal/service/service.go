package service

import (
	"library/internal/repository"
	"library/models"
)

//go:generate mockgen -source=service.go -destination=mock_service.go -package=service
type Authors interface {
	GetAll() ([]models.Author, error)
	Create(author models.Author) error
	GetByID(id int) (models.Author, error)
	Delete(id int) error
	Update(author models.Author) error
}

type Books interface {
	GetAll() ([]models.Book, error)
	Create(book models.Book) error
	GetByID(id int) (models.Book, error)
	Delete(id int) error
	Update(book models.Book) error
	RentBook(userID, bookID int) error
	ReturnBook(userID, bookID int) error
}

type Users interface {
	GetAll() ([]models.User, error)
	Create(user models.User) error
	GetByID(id int) (models.User, error)
	Delete(id int) error
	Update(user models.User) error
}

type Service struct {
	Authors
	Books
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authors: NewAuthorsService(repos.Authors),
		Books:   NewBooksService(repos.Books),
		Users:   NewUsersService(repos.Users),
	}
}
