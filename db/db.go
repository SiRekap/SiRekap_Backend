package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=tungsura port=32768 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
}

func GetDB() *gorm.DB {
	return db
}
