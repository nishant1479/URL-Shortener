package main

import (
	"log"
	"net/http"

	"github.com/nishant1479/URL_Shortener/config"
	"github.com/nishant1479/URL_Shortener/db"
	"github.com/nishant1479/URL_Shortener/handler"
)

func main() {
    // Connect to MongoDB
    mongoCollection := config.ConnectMongo()
    urlRepo := db.NewURLDB(mongoCollection)

    // Connect to Redis
    config.ConnectRedis()

    // Routes
    http.HandleFunc("/shorten", handler.MakeShortenHandler(urlRepo))
    http.HandleFunc("/", handler.MakeRedirectHandler(*urlRepo))


    // TODO: Add GET /{shortKey} redirect handler
    // Example: http.HandleFunc("/", handler.MakeRedirectHandler(urlRepo))

    log.Println("🚀 Server running at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Server error:", err)
    }
}
