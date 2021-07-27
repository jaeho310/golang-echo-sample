package database

import (
	"os"
	"platform-sample/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type SqlStore struct {
}

func (SqlStore) GetDb() *gorm.DB {
	db, err := gorm.Open(os.Getenv("DATASOURCE_DRIVER"), os.Getenv("DATASOURCE_URL"))

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func (SqlStore) GetMockDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./mock.db")

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Card{})
}
