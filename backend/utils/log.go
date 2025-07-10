package utils

import (
	"context"
	"time"
	"net/http"
	"log"
	"github.com/nishant1479/URL_Shortener/config"
	"github.com/nishant1479/URL_Shortener/models"
)

func LogClickAsync(shortKey string, r *http.Request) {
	go func() {
		click := models.ClickEvent{
			ShortKey:  shortKey,
			Timestamp: time.Now(),
			IP:        r.RemoteAddr,
			UserAgent: r.UserAgent(),
		}

		collection := config.MongoClient.Database("urlshortener").Collection("clicks")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		_, err := collection.InsertOne(ctx, click)
		if err != nil {
			log.Println("Failed to log click:", err)
		}
	}()
}
