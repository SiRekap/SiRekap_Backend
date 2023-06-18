package readdb

import (
	"encoding/json"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type FormcImage struct {
	IdImage        int    `json:"id_image" binding:"required" gorm:"primaryKey"`
	JenisImage     int    `json:"jenis_image" binding:"required"`
	NoLembar       int    `json:"no_lembar" binding:"required"`
	IdTps          int    `json:"id_tps" binding:"required"`
	JenisPemilihan int    `json:"jenis_pemilihan" binding:"required"`
	Uuid           string `json:"uuid" binding:"required"`
	NamaFile       string `json:"nama_file" binding:"required"`
	FileHash       string `json:"file_hash" binding:"required"`
	FileSignature  string `json:"file_signature" binding:"required"`
	KodeUsl        string `json:"kode_usl" binding:"required"`
}

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

func init() {
	functions.HTTP("WriteDatabase", WriteDatabase)
}

// WriteDatabase is an HTTP Cloud Function with a request parameter.
func WriteDatabase(w http.ResponseWriter, r *http.Request) {
	Init()

	formcImage := FormcImage{}

	err := json.NewDecoder(r.Body).Decode(&formcImage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	db := GetDB()

	db.Create(&formcImage)
}
