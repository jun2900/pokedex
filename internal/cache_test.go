package internal

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(10 * time.Second)

	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	retrievedVal, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected key to be found, but it was not")
	}
	if string(retrievedVal) != string(val) {
		t.Errorf("expected value to be '%s', but got '%s'", val, retrievedVal)
	}
}

func TestCacheGetNonExistentKey(t *testing.T) {
	cache := NewCache(10 * time.Second)

	key := "non-existent-key"

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key to not be found, but it was")
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(1 * time.Second)

	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	time.Sleep(2 * time.Second)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key to have expired, but it did not")
	}
}

func TestCacheConcurrency(t *testing.T) {
	cache := NewCache(10 * time.Second)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key string, val []byte) {
			defer wg.Done()
			cache.Add(key, val)
		}(fmt.Sprintf("key-%d", i), []byte(fmt.Sprintf("value-%d", i)))
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i)
		val := []byte(fmt.Sprintf("value-%d", i))

		retrievedVal, ok := cache.Get(key)
		if !ok {
			t.Errorf("expected key to be found, but it was not")
		}
		if string(retrievedVal) != string(val) {
			t.Errorf("expected value to be '%s', but got '%s'", val, retrievedVal)
		}
	}
}
