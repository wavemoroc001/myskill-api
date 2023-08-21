package datastore

import (
	"go.mongodb.org/mongo-driver/mongo"
	"myskill-api/model"
)

type User interface {
	GetUserByID(id string) (*model.User, error)
}

type MongoUserDataStore struct {
	col *mongo.Collection
}

func (m *MongoUserDataStore) GetUserByID(id string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewMongoUserDataStore(db *mongo.Database) User {
	return &MongoUserDataStore{col: db.Collection("users")}
}
