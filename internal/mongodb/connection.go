package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseWrapper struct {
	*mongo.Database
}

func Connect() DatabaseWrapper {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	mongodbDatabase := client.Database(os.Getenv("MONGODB_DATABASE_NAME"))
	return DatabaseWrapper{mongodbDatabase}
}

func (db *DatabaseWrapper) Disconnect() {
	err := db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
