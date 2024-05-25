package repositories

import "go.mongodb.org/mongo-driver/mongo"

type mongoRepo struct {
	client *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepo(client *mongo.Client, dbName string, collectionName string) *mongoRepo {
	collection := client.Database(dbName).Collection(collectionName)

	return &mongoRepo{
		client: client,
		collection: collection,
	}
}