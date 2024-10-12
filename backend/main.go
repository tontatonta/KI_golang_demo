package main

import (
	"KI_golang_demo/handlers"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go backend!")
}

func film(w http.ResponseWriter, r *http.Request) {
	filmID := "58611129-2dbc-4a81-a72f-77ddfc1b1b49"
	film, err := handlers.GetFilm(filmID)
	if err != nil {
		http.Error(w, "Error fetching film", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "ID: %s\n", film.ID)
	fmt.Fprintf(w, "Title: %s\n", film.Title)
	fmt.Fprintf(w, "Original Title: %s\n", film.OriginalTitle)
	fmt.Fprintf(w, "Description: %s\n", film.Description)
	fmt.Fprintf(w, "Director: %s\n", film.Director)
	fmt.Fprintf(w, "Producer: %s\n", film.Producer)
	fmt.Fprintf(w, "Release Date: %s\n", film.ReleaseDate)
	fmt.Fprintf(w, "Running Time: %s\n", film.RunningTime)
	fmt.Fprintf(w, "Rotten Tomatoes Score: %s\n", film.RTScore)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/film", film)
	http.ListenAndServe(":8000", nil)

}
