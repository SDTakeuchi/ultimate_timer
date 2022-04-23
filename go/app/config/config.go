package config

import (
	// "ultimate_timer/domain/model"
	"github.com/jinzhu/gorm"
)

// TODO: move to env file
const DbConfig = "host=db port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"

// NewDB DBと接続する
func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", DbConfig)
	if err != nil {
		panic(err.Error())
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// db.AutoMigrate(&model.Preset{})
	db.LogMode(true)

	return db
}
