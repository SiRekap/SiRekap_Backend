package server

import (
	"fmt"
	"sirekap/SiRekap_Backend/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	port := config.GetString("server.port")
	address := fmt.Sprintf(":%s", port)
	r.Run(address)
}
