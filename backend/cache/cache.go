package cache

import (
    "context"
	"github.com/nishant1479/URL_Shortener/config"
    "time"
)

func SetURL(shortKey, originalURL string, ttl time.Duration) error {
    return config.RedisClient.Set(context.TODO(), shortKey, originalURL, ttl).Err()
}

func GetURL(shortKey string) (string, error) {
    return config.RedisClient.Get(context.TODO(), shortKey).Result()
}
