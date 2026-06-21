package storage

import "sc/go13/types"

type MongoDBStorage struct{}

func (m *MongoDBStorage) Get(id int) *types.User {
	return &types.User{
		ID:   id,
		Name: "John Doe",
	}
}
