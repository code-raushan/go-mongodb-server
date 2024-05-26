package services

import (
	"fmt"

	"github.com/code-raushan/go-mongodb-server/types"
)

type InMemoryDB struct {
	store map[string]string
	InMemoryActions
}

type InMemoryActions interface {
	Post(*types.MemoryDataStore) (*types.MemoryDataStore, error)
	Get(string) (*types.MemoryDataStore, error)
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		store: make(map[string]string),
	}
}

func (m *InMemoryDB) Post(data *types.MemoryDataStore) (*types.MemoryDataStore, error){
	m.store[data.Key] = data.Value

	return &types.MemoryDataStore{
		Key: data.Key,
		Value: data.Value,
	}, nil
}

func (m *InMemoryDB) Get(key string) (*types.MemoryDataStore, error){
	data, ok := m.store[key]

	if !ok {
		return nil, fmt.Errorf("Key not found")
	}

	return &types.MemoryDataStore{
		Key: key,
		Value: data,
	}, nil
}