package db

import (
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "sirekap/SiRekap_Backend/config"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	// c := config.GetConfig()

	dsn := "host=localhost user=postgres password=postgrespw dbname=rekapitulasi port=32768 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db = dynamodb.New(session.New(&aws.Config{
	// 	Region:      aws.String(c.GetString("db.region")),
	// 	Credentials: credentials.NewEnvCredentials(),
	// 	Endpoint:    aws.String(c.GetString("db.endpoint")),
	// 	DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),
	// }))
}

func GetDB() *gorm.DB {
	return db
}
