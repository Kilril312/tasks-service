package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Kilril312/tasks-service/internal/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&task.Task{}); err != nil {
		log.Printf("Migration error: %v", err)
		db.Migrator().DropTable(&task.Task{})
		db.AutoMigrate(&task.Task{})
	}

	return db, nil
}
