package repository

import (
	"gorm.io/gorm"
	"library/models"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("RentedBooks").Find(&users).Error
	return users, err
}

func (r *UserPostgres) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserPostgres) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("RentedBooks").First(&user, id).Error
	return user, err
}

func (r *UserPostgres) Delete(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserPostgres) Update(user models.User) error {
	return r.db.Save(&user).Error
}
