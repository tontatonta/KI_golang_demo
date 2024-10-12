package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// ハンドラーを登録
	http.HandleFunc("/", HelloHandler)

	// サーバーのポート番号
	port := "8080"

	// HTTPS URLを表示する
	fmt.Printf("サーバーを以下のURLで起動中: http://localhost:%s\n", port)

	// HTTPサーバーを起動
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("サーバーの起動に失敗:", err)
	}
}
