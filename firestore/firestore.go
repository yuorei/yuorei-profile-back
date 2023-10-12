package fiestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	var app *firebase.App
	var err error

	// Firestoreの初期化
	if "prod" == os.Getenv("PROD") {
		conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
		app, err = firebase.NewApp(ctx, conf)
		if err != nil {
			log.Fatalf("Failed to initialize Firestore: %v", err)
		}
	} else {
		opt := option.WithCredentialsFile("path/to/serviceAccount.json")
		app, err = firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Fatalf("Failed to initialize Firestore: %v", err)
		}
	}

	// Firestoreのクライアントを取得
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return client, err
}
