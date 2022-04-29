package config

import (
	// "ultimate_timer/domain/model"
	"fmt"
	"os"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)

func NewDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// db.AutoMigrate(&model.Preset{})
	db.Logger = db.Logger.LogMode(logger.Info)

	return db
}
