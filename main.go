package main

import (
	fiestore "blog/firestore"
	"blog/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()
	client, err := fiestore.NewFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}
	h := handler.NewHandler(client)

	r.HandleFunc("/blog", h.GetBlogs).Methods("GET")
	r.HandleFunc("/blog/{id}", h.GetBlogByTitle).Methods("GET")

	http.Handle("/", r)
	fmt.Println("起動" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
