package notifications

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"sirekap/SiRekap_Backend/forms"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var app *firebase.App

func Init() {

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	app = firebaseApp
}

func SendToToken(registrationToken string, form forms.FormcScanResponse) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
		return err
	}

	formJson, err := json.Marshal(form)
	if err != nil {
		fmt.Println(err)
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Form C Scan Response Result",
			Body:  string(formJson),
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	fmt.Println("Successfully sent message:", response)
	return nil
}
