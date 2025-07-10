package handler

import (
    "github.com/nishant1479/URL_Shortener/service"
    "github.com/nishant1479/URL_Shortener/db"
    "log"
    "net/http"
    "strings"
)

func MakeRedirectHandler(repo db.URLDB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract key from path like /abc123
        path := strings.TrimPrefix(r.URL.Path, "/")
        if path == "" {
            http.Error(w, "Missing short URL key", http.StatusBadRequest)
            return
        }

        // Use service to resolve the key
        originalURL, err := service.ResolveURL(path, repo)
        if err != nil {
            log.Println("Resolve error:", err)
            http.Error(w, "URL not found or expired", http.StatusNotFound)
            return
        }

        http.Redirect(w, r, originalURL, http.StatusFound)
    }
}
