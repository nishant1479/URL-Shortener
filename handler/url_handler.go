package handler

import (
    "encoding/json"
    "net/http"
	"github.com/nishant1479/URL_Shortener/service"
	"github.com/nishant1479/URL_Shortener/db"
)

type ShortenRequest struct {
    OriginalURL     string `json:"original_url"`
    ValidForMinutes int    `json:"valid_for_minutes"`
}

type ShortenResponse struct {
    ShortURL string `json:"short_url"`
}

func MakeShortenHandler(repo *db.URLDB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req ShortenRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil || req.OriginalURL == "" {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        // Default expiry: 1440 minutes (1 day)
        if req.ValidForMinutes == 0 {
            req.ValidForMinutes = 1440
        }

        shortKey, err := service.ShortenURL(req.OriginalURL, req.ValidForMinutes, *repo)
        if err != nil {
            http.Error(w, "Error creating short URL: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // Build full short URL
        shortURL := "http://localhost:8080/" + shortKey

        res := ShortenResponse{ShortURL: shortURL}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(res)
    }
}
