package main

import (
	"KI_golang_demo/handlers"
	"encoding/json"
	"net/http"
)

// CORS設定のためのミドルウェア
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go backend!"))
}

func film(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	filmID := "58611129-2dbc-4a81-a72f-77ddfc1b1b49"
	film, err := handlers.GetFilm(filmID)
	if err != nil {
		http.Error(w, "Error fetching film", http.StatusInternalServerError)
		return
	}

	// JSONとしてレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/film", film)
	http.ListenAndServe(":8000", nil)
}
