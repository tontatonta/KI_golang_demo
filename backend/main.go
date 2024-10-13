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

	// クエリパラメーターから映画IDを取得
	filmID := r.URL.Query().Get("id")
	if filmID == "" {
		http.Error(w, "Film ID is required", http.StatusBadRequest)
		return
	}

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
