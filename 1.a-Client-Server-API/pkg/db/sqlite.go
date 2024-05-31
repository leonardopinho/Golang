package db

import (
	"context"
	"github.com/leonardopinho/GoLang/1.a-Client-Server-API/pkg/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func InitDb() error {
	var err error

	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&models.USDBRL{})
	if err != nil {
		return err
	}

	return nil
}

func SaveUSDBRL(data *models.USDBRL) (*models.USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&data)
		if result.Error != nil {
			log.Fatal(result.Error)
			return result.Error
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}

func DeleteUSDBRL(data *models.USDBRL) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&data)
		if result.Error != nil {
			log.Fatal(result.Error)
			return result.Error
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
