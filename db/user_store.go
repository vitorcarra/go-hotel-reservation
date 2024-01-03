package db

import (
	"context"

	"github.com/vitorcarra/go-hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	dbName     string
	collection *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client, dbName string, coll string) *MongoUserStore {
	return &MongoUserStore{
		client:     c,
		collection: c.Database(dbName).Collection(coll),
		dbName:     dbName,
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	// validate the ID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
