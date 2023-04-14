package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() {
	dsn := "host=db.ifylprocwbwlhjupjbrl.supabase.co user=postgres password=JhfXT5UlxKOTPKXF dbname=postgres port=5432 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
}

func GetDB() *gorm.DB {
	return db
}
