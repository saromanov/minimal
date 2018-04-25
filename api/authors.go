package api

import (
	"encoding/json"
	"net/http"
)

// Author provides definition for author api struct
type Author struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	JoinedAt    string `json:"joinedAt"`
}

// CreateAuthor provides create of the new author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author Author
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
