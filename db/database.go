package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase(DBdriver string, DBSource string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(DBSource), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}
