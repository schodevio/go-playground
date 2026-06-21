package storage

import "sc/go13/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		ID:   id,
		Name: "John Doe",
	}
}
