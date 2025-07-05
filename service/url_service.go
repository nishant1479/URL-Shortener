package service

import (
	"context"
	"fmt"
	"time"
	"github.com/nishant1479/URL_Shortener/models"
	"github.com/nishant1479/URL_Shortener/db"
	"github.com/nishant1479/URL_Shortener/utils"
)


type URLService interface {
    ShortenURL(originalURL string, validMinutes int) (string, error)
    ResolveURL(shortKey string) (string, error)
}


func ShortenURL(originalURL string, validMinutes int, repo db.URLDB) (string, error) {
    // Step 1: Validate the URL
    if !utils.IsValidURL(originalURL) {
        return "", fmt.Errorf("invalid URL")
    }

    // Step 2: Generate a short key
    shortKey := utils.GenerateShortKey(6)

    // Step 3: Create document
    now := time.Now()
    expiration := now.Add(time.Duration(validMinutes) * time.Minute)

    urlDoc := models.URL{
        ID:          shortKey,
        OriginalURL: originalURL,
        CreatedAt:   now,
        Expiration:  expiration,
    }

    // Step 4: Save to MongoDB
    err := repo.InsertURL(context.TODO(), urlDoc)
    if err != nil {
        return "", err
    }

    return shortKey, nil
}
func ResolveURL(shortKey string, repo db.URLDB) (string, error) {
    urlDoc, err := repo.FindByShortKey(context.TODO(), shortKey)
    if err != nil {
        return "", err
    }

    // Check for expiration
    if time.Now().After(urlDoc.Expiration) {
        return "", fmt.Errorf("link expired")
    }

    return urlDoc.OriginalURL, nil
}
