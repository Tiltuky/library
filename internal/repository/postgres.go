package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"library/models"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Author{}, &models.Book{}, &models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
