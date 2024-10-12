package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FilmはAPIレスポンスの構造を定義します
type Film struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Description   string `json:"description"`
	Director      string `json:"director"`
	Producer      string `json:"producer"`
	ReleaseDate   string `json:"release_date"`
	RunningTime   string `json:"running_time"`
	RTScore       string `json:"rt_score"`
}

// GetFilmは指定されたIDの映画を取得します
func GetFilm(filmID string) (*Film, error) {
	url := fmt.Sprintf("https://ghibliapi.vercel.app/films/%s", filmID)

	// GETリクエストを送信
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// レスポンスボディを読み取る
	body, err := io.ReadAll(response.Body) // 修正: ioutil.ReadAllからio.ReadAllに変更
	if err != nil {
		return nil, err
	}

	// JSONを構造体にデコード
	var film Film
	err = json.Unmarshal(body, &film)
	if err != nil {
		return nil, err
	}

	return &film, nil
}
