package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"sirekap/SiRekap_Backend/config"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/server"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
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

	return nil
}

func main() {
	fc.StartHttp(HandleHttpRequest)
}
