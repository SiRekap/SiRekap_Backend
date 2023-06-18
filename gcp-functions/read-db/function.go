package readdb

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Tps struct {
	IdTps          int    `json:"id_tps" binding:"required" gorm:"primaryKey"`
	Alamat         string `json:"alamat" binding:"required"`
	Koordinat      string `json:"koordinat" binding:"required"`
	Tipe           string `json:"tipe" binding:"required"`
	IdWilayahDasar int    `json:"id_wilayah_dasar" binding:"required"`
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
	functions.HTTP("ReadDatabase", ReadDatabase)
}

// ReadDatabase is an HTTP Cloud Function with a request parameter.
func ReadDatabase(w http.ResponseWriter, r *http.Request) {
	Init()

	_, err := GetTpsById(1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprintf(w, "tps")
}

func GetTpsById(idTps int) (Tps, error) {

	db := GetDB()

	tps := Tps{}

	result := db.Where("id_tps = ?", idTps).First(&tps)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Tps{}, result.Error
	} else if result.Error != nil {
		return Tps{}, result.Error
	}

	return tps, nil
}
