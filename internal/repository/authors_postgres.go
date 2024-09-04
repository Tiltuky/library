package repository

import (
	"gorm.io/gorm"
	"library/models"
)

type AuthorPostgres struct {
	db *gorm.DB
}

func NewAuthorPostgres(db *gorm.DB) *AuthorPostgres {
	return &AuthorPostgres{db: db}
}

func (r *AuthorPostgres) GetAll() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *AuthorPostgres) Create(author models.Author) error {
	return r.db.Create(&author).Error
}

func (r *AuthorPostgres) GetByID(id int) (models.Author, error) {
	var author models.Author
	err := r.db.First(&author, id).Error
	return author, err
}

func (r *AuthorPostgres) Delete(id int) error {
	return r.db.Delete(&models.Author{}, id).Error
}

func (r *AuthorPostgres) Update(author models.Author) error {
	return r.db.Save(&author).Error
}
