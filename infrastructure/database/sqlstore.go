package database

import (
	"os"
	"platform-sample/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type SqlStore struct {
}

func (SqlStore) Getdb() *gorm.DB {
	return connection()
}

func connection() *gorm.DB {
	db, err := gorm.Open(os.Getenv("DATASOURCE_DRIVER"), os.Getenv("DATASOURCE_URL"))

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
