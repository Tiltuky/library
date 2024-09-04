package models

import "time"

type Author struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Books []Book
}

type Book struct {
	ID          int       `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	AuthorID    int       `gorm:"not null"`
	PublishedAt time.Time `gorm:"not null"`
	ISBN        string    `gorm:"unique;not null"`
	Author      Author
	RentedBooks []RentedBook
}

type User struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	RentedBooks []RentedBook
}

type RentedBook struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	BookID     int       `gorm:"not null"`
	RentedAt   time.Time `gorm:"not null"`
	ReturnedAt *time.Time
	Book       Book
}
