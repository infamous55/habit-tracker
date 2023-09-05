package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DatabaseWrapper) CreateIndex(
	collectionName string,
	field string,
	unique bool,
) error {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		return err
	}

	return nil
}
