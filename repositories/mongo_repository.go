package repositories

import (
	"context"
	"log"

	"github.com/code-raushan/go-mongodb-server/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepo(client *mongo.Client, dbName string, collectionName string) *MongoRepo {
	collection := client.Database(dbName).Collection(collectionName)

	return &MongoRepo{
		client:     client,
		collection: collection,
	}
}

func (m *MongoRepo) Fetch(filter types.FilterOptions) []types.FetchResponse {
	ctx := context.Background()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$and": []bson.M{
					{"createdAt": bson.M{"$gte": filter.StartDate}},
					{"createdAt": bson.M{"$lte": filter.EndDate}},
				},
			},
		},
		{
			"$addFields": bson.M{
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"$and": []bson.M{
					{"totalCount": bson.M{"$gte": filter.MinCount}},
					{"totalCount": bson.M{"$lt": filter.MaxCount}},
				},
			},
		},
		{"$project": bson.M{
			"key":        1,
			"createdAt":  1,
			"totalCount": 1,
		}},
	}

	cur, err := m.collection.Aggregate(ctx, pipeline)

	if err != nil {
		log.Fatalf("Error while fetching data records %v", err)
	}

	defer cur.Close(ctx)

	var results []types.FetchResponse
	if err := cur.All(ctx, &results); err != nil {
		log.Fatalf("Error while storing data records into response: %v", err)
	}

	return results
}
