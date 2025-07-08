package middleware

import (
    "net/http"
    "os"
)

func APIKeyAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        key := r.Header.Get("X-API-Key")
        if key == "" || key != os.Getenv("API_KEY") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}