package repository

import (
	"gorm.io/gorm"
	"library/models"
)

//go:generate mockgen -source=repository.go -destination=mock_repository.go -package=repository
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

type Repository struct {
	Authors
	Books
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authors: NewAuthorPostgres(db),
		Books:   NewBookPostgres(db),
		Users:   NewUserPostgres(db),
	}
}
