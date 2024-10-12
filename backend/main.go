package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go backend!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :8000")
    http.ListenAndServe(":8000", nil)
}
