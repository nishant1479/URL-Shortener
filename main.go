package main

import (
    "net/http"
	"github.com/nishant1479/URL_Shortener/db"
	"github.com/nishant1479/URL_Shortener/handler"
	"github.com/nishant1479/URL_Shortener/config"
)

func main() {
    collection := config.ConnectMongo()
    repo := db.NewURLDB(collection)

    http.HandleFunc("/shorten", handler.MakeShortenHandler(repo))

    // TODO: Add GET /{shortKey} handler later
    http.ListenAndServe(":8080", nil)
}
