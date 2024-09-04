package data

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"library/models"
	"time"
)

func InitData(db *gorm.DB) error {
	// Initialize authors
	if err := initAuthors(db); err != nil {
		return err
	}
	// Initialize books
	if err := initBooks(db); err != nil {
		return err
	}
	// Initialize users
	if err := initUsers(db); err != nil {
		return err
	}

	return nil
}

func initAuthors(db *gorm.DB) error {
	var count int64
	db.Model(&models.Author{}).Count(&count)
	if count == 0 {
		authors := make([]models.Author, 10)
		for i := 0; i < 10; i++ {
			authors[i] = models.Author{
				Name: gofakeit.Name(),
			}
		}
		if err := db.Create(&authors).Error; err != nil {
			return fmt.Errorf("failed to create authors: %w", err)
		}
	}
	return nil
}

func initBooks(db *gorm.DB) error {
	var count int64
	db.Model(&models.Book{}).Count(&count)
	if count == 0 {
		var authors []models.Author
		db.Find(&authors)
		if len(authors) == 0 {
			return fmt.Errorf("no authors found to assign books")
		}

		books := make([]models.Book, 100)
		for i := 0; i < 100; i++ {
			books[i] = models.Book{
				Title:       gofakeit.BookTitle(),
				AuthorID:    authors[gofakeit.Number(0, len(authors)-1)].ID,
				PublishedAt: gofakeit.DateRange(time.Now().AddDate(-10, 0, 0), time.Now()),
				ISBN:        generateISBN(),
			}
		}
		if err := db.Create(&books).Error; err != nil {
			return fmt.Errorf("failed to create books: %w", err)
		}
	}
	return nil
}

func initUsers(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		users := make([]models.User, 50)
		for i := 0; i < 50; i++ {
			users[i] = models.User{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			}
		}
		if err := db.Create(&users).Error; err != nil {
			return fmt.Errorf("failed to create users: %w", err)
		}
	}
	return nil
}

func generateISBN() string {
	return fmt.Sprintf("%d-%d-%d-%d-%d",
		gofakeit.Number(100, 999),
		gofakeit.Number(1000, 9999),
		gofakeit.Number(100, 999),
		gofakeit.Number(10, 99),
		gofakeit.Number(1000, 9999),
	)
}
