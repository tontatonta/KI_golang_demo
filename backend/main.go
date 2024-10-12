package main

import (
    "fmt"
    "net/http"
    "github.com/tontatonta/KI_golang_demo/backend/src/film" // filmパッケージをインポート
)

// ハンドラー関数でfilmパッケージを使用
func handler(w http.ResponseWriter, r *http.Request) {
    // filmパッケージの関数を使って映画のタイトルを取得
    title := film.GetTitle()
    fmt.Fprintf(w, "Hello from Go backend!! Here is a movie recommendation: %s", title)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :8000")
    http.ListenAndServe(":8000", nil)
}

