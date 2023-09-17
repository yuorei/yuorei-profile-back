package fiestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirestoreClient() (*firestore.Client, error) {
	// Firestoreの初期化
	opt := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	// Firestoreのクライアントを取得
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return client, err
}
