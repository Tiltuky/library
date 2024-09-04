package repository

import (
	"fmt"
	"gorm.io/gorm"
	"library/models"
	"time"
)

type BookPostgres struct {
	db *gorm.DB
}

func NewBookPostgres(db *gorm.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Preload("Author").Find(&books).Error
	return books, err
}

func (r *BookPostgres) Create(book models.Book) error {
	return r.db.Create(&book).Error
}

func (r *BookPostgres) GetByID(id int) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("Author").First(&book, id).Error
	return book, err
}

func (r *BookPostgres) Delete(id int) error {
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *BookPostgres) Update(book models.Book) error {
	return r.db.Save(&book).Error
}

func (r *BookPostgres) RentBook(userID, bookID int) error {
	var rentedBook models.RentedBook
	if err := r.db.Where("book_id = ? AND returned_at IS NULL", bookID).First(&rentedBook).Error; err == nil {
		return fmt.Errorf("book is already rented")
	}

	rentedBook = models.RentedBook{
		UserID:   userID,
		BookID:   bookID,
		RentedAt: time.Now(),
	}

	return r.db.Create(&rentedBook).Error
}

func (r *BookPostgres) ReturnBook(userID, bookID int) error {
	var rentedBook models.RentedBook
	if err := r.db.Where("user_id = ? AND book_id = ? AND returned_at IS NULL", userID, bookID).First(&rentedBook).Error; err != nil {
		return fmt.Errorf("book is not rented by this user")
	}

	now := time.Now()
	rentedBook.ReturnedAt = &now
	return r.db.Save(&rentedBook).Error
}
