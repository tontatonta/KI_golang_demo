package film

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Film represents a structure for the Ghibli film data
type Film struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Description   string `json:"description"`
	Director      string `json:"director"`
	Producer      string `json:"producer"`
	ReleaseDate   string `json:"release_date"`
	RunningTime   string `json:"running_time"`
	RtScore       string `json:"rt_score"`
}

// GetFilms retrieves the list of Ghibli films from the external API
func GetFilms() ([]Film, error) {
	resp, err := http.Get("https://ghibliapi.vercel.app/films")
	if err != nil {
		return nil, fmt.Errorf("failed to get films: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Replace ioutil.ReadAll with io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var films []Film
	if err := json.Unmarshal(body, &films); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return films, nil
}
