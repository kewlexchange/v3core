package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"core/models"
	seed "core/seeders"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // Global değişken olarak veritabanı bağlantısı

func InitDB() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is required")
	}

	errorOnlyLogger := logger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Error, // sadece Error
			IgnoreRecordNotFoundError: true,         // record not found'u loglama
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: errorOnlyLogger})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		// Hata işleme
	}

	sqlDB.SetMaxIdleConns(10)           // Boşta bekleyen bağlantıların maksimum sayısı
	sqlDB.SetMaxOpenConns(0)            // Aynı anda açık olabilecek maksimum bağlantı sayısı
	sqlDB.SetConnMaxLifetime(time.Hour) // Bağlantının yeniden kullanılabilir olacağı maksimum süre

	DB = db
	return nil
}

func Migrate(db *gorm.DB) error {
	fmt.Println("Migration:Begin")
	//db.Logger = db.Logger.LogMode(logger.Silent)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	err := db.AutoMigrate(
		&models.Currency{},
		&models.Exchange{},
		&models.Pair{},
	)

	return err
}

func Seed(db *gorm.DB) error {
	fmt.Println("Seed Begin")
	seed.Seed(db)
	fmt.Println("Seed End")
	return nil
}
