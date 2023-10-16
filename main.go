package main

import (
	fiestore "blog/firestore"
	"blog/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	r := mux.NewRouter()
	client, err := fiestore.NewFirestoreClient()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}
	h := handler.NewHandler(client)

	r.HandleFunc("/blog", h.GetBlogs).Methods("GET")
	r.HandleFunc("/blog/{id}", h.GetBlogByID).Methods("GET")
	http.Handle("/", r)
	log.Println("起動" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
