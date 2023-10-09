package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

type Handler struct {
	client *firestore.Client
}

func NewHandler(client *firestore.Client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) GetBlogs(w http.ResponseWriter, r *http.Request) {
	// CORS対応: 必要なヘッダーを追加
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Firestoreからデータを取得
	collection := h.client.Collection("blog")
	documents, err := collection.Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("Failed to retrieve documents: %v", err)
	}

	// レスポンスのJSONを作成
	var results []map[string]any
	for _, doc := range documents {
		data := doc.Data()
		results = append(results, data)
	}

	// JSONをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (h *Handler) GetBlogByTitle(w http.ResponseWriter, r *http.Request) {
	// CORS対応: 必要なヘッダーを追加
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// クエリパラメーターからタイトルを取得
	params := mux.Vars(r)
	id := params["id"]

	dsnap, err := h.client.Collection("blog").Doc(id).Get(context.Background())
	if err != nil {
		return
	}
	result := dsnap.Data()

	// JSONをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
