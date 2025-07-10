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


// MongoClient is the shared global MongoDB client
var MongoClient *mongo.Client

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

// InitMongo initializes the MongoDB client and assigns it to MongoClient
func InitMongo() {
    LoadEnv()

    uri := os.Getenv("MONGODB_URI")
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    MongoClient, err = mongo.Connect(ctx, opts)
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    if err := MongoClient.Ping(ctx, nil); err != nil {
        log.Fatalf("MongoDB ping error: %v", err)
    }

    log.Println("✅ Connected to MongoDB")
}

// GetCollection returns a MongoDB collection from the specified database

func GetCollection(dbName, collectionName string) *mongo.Collection {
    return MongoClient.Database(dbName).Collection(collectionName)
}

// the reason we change the code of this as now we have to get the data of the number of clicks, which we can get by calling the whole collection
