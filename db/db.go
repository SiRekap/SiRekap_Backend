package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() {
	dsn := "host=db.zoeczoxdgyuhezioxvcy.supabase.co user=postgres password=Pqq2#Jd/ZgDL*5D dbname=postgres port=5432 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
}

func GetDB() *gorm.DB {
	return db
}
