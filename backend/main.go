package main

import (
	"encoding/json"
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

	// ✅ Start cleanup job
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for {
			<-ticker.C
			utils.RemoveExpiredLinks()
		}
	}()

	// ✅ Routes
	http.HandleFunc("/shorten", middleware.APIKeyAuth(handler.MakeShortenHandler(urlRepo)))
	http.HandleFunc("/", handler.MakeRedirectHandler(*urlRepo))
	http.HandleFunc("/api/data", handleData) // <-- this was unreachable before

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

// ✅ React frontend can fetch this: http://localhost:8080/api/data
func handleData(w http.ResponseWriter, r *http.Request) {
	// ✅ Allow frontend from different port
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "Hello from Go backend!",
		"status":  "success",
	}

	json.NewEncoder(w).Encode(response)
}
