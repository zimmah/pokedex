package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entry 	map[string] cacheEntry
	mux 	*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entry: 	make(map[string]cacheEntry),
		mux:	&sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (cache Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()

	cache.entry[key] = cacheEntry{
		createdAt:	time.Now().UTC(),
		val: 		val,
	}
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	entry, exists := cache.entry[key]
	if !exists { return nil, false }
	return entry.val, exists
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache Cache) reap(now time.Time, last time.Duration) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	for key, value := range cache.entry {
		if value.createdAt.Before(now.Add(-last)) {
			delete(cache.entry, key)
		}
	}
}