package utils

import (
    "context"
    "log"
    "os"
    "time"

	"github.com/nishant1479/URL_Shortener/config"
    "go.mongodb.org/mongo-driver/bson"
)

func RemoveExpiredLinks() {
    collection := config.GetCollection(os.Getenv("MONGO_DB_NAME"), os.Getenv("MONGO_COLLECTION_NAME"))

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{
        "expiration": bson.M{"$lt": time.Now()},
    }

    result, err := collection.DeleteMany(ctx, filter)
    if err != nil {
        log.Println("⚠️ Cleanup error:", err)
        return
    }

    log.Printf("🧹 Cleanup complete. Removed %d expired links.\n", result.DeletedCount)
}
