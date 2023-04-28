package main

import (
	"flag"
	"fmt"
	"os"
	"sirekap/SiRekap_Backend/config"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	// migrations.Migrate()
	server.Init()

	// controllers.GeneratePdfAndSendToBucket(
	// 	"https://wallpapers.com/images/featured/2ygv7ssy2k0lxlzu.jpg",
	// 	"https://wallpapers.com/images/featured/2ygv7ssy2k0lxlzu.jpg",
	// 	"https://wallpapers.com/images/featured/2ygv7ssy2k0lxlzu.jpg",
	// )
}
