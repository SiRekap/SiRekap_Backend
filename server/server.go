package server

import (
	"sirekap/SiRekap_Backend/config"
	"fmt"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	host := config.GetString("server.host")
	port := config.GetString("server.port")
	address := fmt.Sprintf("%s:%s", host, port)
	r.Run(address)
}
