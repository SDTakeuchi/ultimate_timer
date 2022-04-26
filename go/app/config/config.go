package config

import (
	// "ultimate_timer/domain/model"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SSLMODE"),
	)
	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		panic(err.Error())
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// db.AutoMigrate(&model.Preset{})
	db.LogMode(true)

	return db
}
