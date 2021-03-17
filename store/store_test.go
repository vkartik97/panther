package store

import (
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	dataStore := store{
		dataMap: make(map[string]string),
	}
	for i:=0 ; i <= b.N ; i++ {
		go dataStore.Put("a", "123")
	}
}

func BenchmarkRW(b *testing.B) {
	dataStore := store{
		dataMap: make(map[string]string),
	}
	for i:=0 ; i <= b.N ; i++ {
		dataStore.Put("a", "123")
		go dataStore.Get("a")
	}
}
