package config

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Collection {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    uri := os.Getenv("MONGODB_URI")
    dbName := os.Getenv("MONGO_DB_NAME")
    collectionName := os.Getenv("MONGO_COLLECTION_NAME")

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, opts)
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatalf("MongoDB ping error: %v", err)
    }

    log.Println("The connection to MongoDB has been established")
    return client.Database(dbName).Collection(collectionName)
}
