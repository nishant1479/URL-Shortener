package service

import (
    "context"
	"fmt"
	"time"
    "github.com/nishant1479/URL_Shortener/cache"
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
    resultCh := make(chan string)
    errCh := make(chan error)

    //  Redis goroutine
    go func() {
        url, err := cache.GetURL(shortKey)
        if err == nil && url != "" {
            resultCh <- url
        } else {
            errCh <- err
        }
    }()

    //  MongoDB goroutine
    go func() {
        doc, err := repo.FindByShortKey(context.TODO(), shortKey)
        if err != nil {
            errCh <- err
            return
        }

        // Check expiration
        if time.Now().After(doc.Expiration) {
            errCh <- fmt.Errorf("URL expired")
            return
        }

        // Cache in Redis for future
        go cache.SetURL(shortKey, doc.OriginalURL, time.Until(doc.Expiration))

        resultCh <- doc.OriginalURL
    }()

    //  Whichever wins the race returns first
    select {
    case url := <-resultCh:
        return url, nil
    case err := <-errCh:
        return "", err
    case <-time.After(2 * time.Second): // timeout safety
        return "", fmt.Errorf("timeout resolving URL")
    }
}
