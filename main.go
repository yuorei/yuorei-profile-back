package main

import (
	fiestore "blog/firestore"
	"blog/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client, err := fiestore.NewFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}
	h := handler.NewHandler(client)

	http.HandleFunc("/blog", h.GetBlogs)
	fmt.Println("起動")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
