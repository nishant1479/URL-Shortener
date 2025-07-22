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
	// Step 1: Initialize MongoDB connection
	config.InitMongo() //  Add this FIRST

	// Step 2: Get collection AFTER InitMongo
	mongoCollection := config.GetCollection("urlshortener", "urls")
	urlRepo := db.NewURLDB(mongoCollection)

	// Step 3: Connect to Redis
	config.ConnectRedis()

	// Step 4: Start cleanup job
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for {
			<-ticker.C
			utils.RemoveExpiredLinks()
		}
	}()

	// Step 5: Define Routes
	http.HandleFunc("/shorten", middleware.APIKeyAuth(handler.MakeShortenHandler(urlRepo)))
	http.HandleFunc("/api/data", handleData)

	// Serve a welcome message at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handler.MakeRedirectHandler(*urlRepo)(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to the URL Shortener API! Use /shorten to create short URLs.",
		})
	})

	// Catch-all handler for undefined routes (not strictly needed, as above closure handles all / paths)
	// If you want to handle truly undefined routes, use a custom mux instead of http.DefaultServeMux

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(" Server error:", err)
	}
}

// React frontend can fetch this: http://localhost:8080/api/data
func handleData(w http.ResponseWriter, r *http.Request) {
	// Allow frontend from different port
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "Hello from Go backend!",
		"status":  "success",
	}

	json.NewEncoder(w).Encode(response)
}
