package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nishant1479/URL_Shortener/config"
	"github.com/nishant1479/URL_Shortener/db"
	"github.com/nishant1479/URL_Shortener/handler"
	"github.com/nishant1479/URL_Shortener/middleware"
	"github.com/nishant1479/URL_Shortener/utils"
)

func main() {
	mongoCollection := config.GetCollection("urlshortener", "urls")
	urlRepo := db.NewURLDB(mongoCollection)

	config.ConnectRedis()

	// Routes
	http.HandleFunc("/shorten", middleware.APIKeyAuth(handler.MakeShortenHandler(urlRepo)))
	http.HandleFunc("/", handler.MakeRedirectHandler(*urlRepo))

	// TODO: Add GET /{shortKey} redirect handler
	// Example: http.HandleFunc("/", handler.MakeRedirectHandler(urlRepo))

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
    go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for {
			<-ticker.C
			utils.RemoveExpiredLinks()
		}
	}()
}
