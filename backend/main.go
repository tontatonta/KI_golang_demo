package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/tontatonta/KI_golang_demo/contents/film" // filmパッケージをインポート
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	films, err := film.GetFilms() // filmパッケージの関数を呼び出す
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}

func main() {
	// ハンドラーを登録
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/films", filmsHandler)

	// サーバーのポート番号
	port := "8080"

	// HTTPS URLを表示する
	fmt.Printf("サーバーを以下のURLで起動中: http://localhost:%s\n", port)

	// HTTPサーバーを起動
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("サーバーの起動に失敗:", err)
	}
}
