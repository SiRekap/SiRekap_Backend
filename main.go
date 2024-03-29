package main

import (
	"flag"
	"fmt"
	"os"
	"sirekap/SiRekap_Backend/config"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/notifications"
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
	notifications.Init()
	// migrations.Migrate()
	server.Init()

	// controllers.SendFormcResultStreamProcessingRequestTest()

	// controllers.GeneratePdfAndSendToBucket(
	// 	"http://i.ibb.co/0YdPt2F/Whats-App-Image-2023-04-27-at-6-30-05-AM.jpg",
	// 	"http://storage.googleapis.com/staging-sirekap-form/1682425578819.png",
	// 	"http://storage.googleapis.com/staging-sirekap-form/1682425575786.png",
	// 	"204.pdf",
	// )

	// res, err := controllers.SendFormcImageVisionRequest(forms.FormcImageRawResponse{
	// 	FormcImageRaw: forms.FormcImageRaw{
	// 		PayloadList:  []string{"https://i.ibb.co/0YdPt2F/Whats-App-Image-2023-04-27-at-6-30-05-AM.jpg", "https://storage.googleapis.com/staging-sirekap-form/1682425578819.png", "https://storage.googleapis.com/staging-sirekap-form/1682425575786.png"},
	// 		IdPaslonList: []int{1, 2, 3},
	// 	},
	// 	IdImageList: []int{1, 2, 3},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(res)
}
