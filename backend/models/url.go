package models

import "time"

type URL struct {
    ID          string    `bson:"_id" json:"short_key"`          // Short key (e.g., "abc123")
    OriginalURL string    `bson:"original_url" json:"original_url"` // Long URL
    CreatedAt   time.Time `bson:"created_at" json:"created_at"`     // When it was created
    Expiration  time.Time `bson:"expiration" json:"expiration"`     // When it will expires
}
