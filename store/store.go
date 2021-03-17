package store

import (
	"sync"
)

type store struct {
	sync.RWMutex
	dataMap map[string] string
}

func NewDataStore() *store {
	dataStore := store{
		dataMap: make(map[string]string),
	}
	return &dataStore
}

func (dataStore *store) Get(key string) string {
	dataStore.RLock()
	defer dataStore.RUnlock()
	return dataStore.dataMap[key]
}

func (dataStore *store) Put(key, value string) {
	dataStore.Lock()
	defer dataStore.Unlock()
	dataStore.dataMap[key] = value
}
