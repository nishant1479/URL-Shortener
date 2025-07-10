package db
import (
	"context"

	"github.com/nishant1479/URL_Shortener/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLDB struct {
    Collection *mongo.Collection
}

func NewURLDB(c *mongo.Collection) *URLDB {
	return &URLDB{
		Collection: c,
	}
}

func (r *URLDB) InsertURL(ctx context.Context, url models.URL) error {
    _, err := r.Collection.InsertOne(ctx, url)
    return err
}

func (r *URLDB) FindByShortKey(ctx context.Context, key string) (models.URL, error) {
    var url models.URL
    err := r.Collection.FindOne(ctx, bson.M{"_id": key}).Decode(&url)
    return url, err
}

