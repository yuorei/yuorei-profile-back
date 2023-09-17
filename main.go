package main

import (
	fiestore "blog/firestore"
	"blog/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	client, err := fiestore.NewFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}
	h := handler.NewHandler(client)

	r.HandleFunc("/blog", h.GetBlogs).Methods("GET")
	r.HandleFunc("/blog/{title}", h.GetBlogByTitle).Methods("GET")
	fmt.Println("起動")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
