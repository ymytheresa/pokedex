package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []string
}

type Cache struct {
	cacheMemory map[string]cacheEntry
	mu          sync.Mutex
	interval    time.Duration
}

func (c Cache) readLoop(interval time.Duration) {
	c.interval = interval * time.Second
	ticker := time.Tick(interval * time.Second)
	go func() {
		for {
			select {
			case <-ticker:
				cleanUpCache(c)
			}
		}
	}()
}

func cleanUpCache(c Cache) {
	c.mu.Lock()
	defer c.mu.Unlock()

	mem := c.cacheMemory
	for key, valueEntry := range mem {
		if key == "localPokedex" {
			continue //do not clear record of caught pet
		}
		expiryTime := time.Now().Add(-c.interval)
		if valueEntry.createdAt.Before(expiryTime) {
			delete(mem, key)
		}
	}
}

func (c Cache) Add(key string, entryVal []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMemory[key] = cacheEntry{
		createdAt: time.Now(),
		val:       entryVal,
	}
}

func (c Cache) Get(key string) ([]string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.cacheMemory[key]
	if ok {
		return value.val, true
	}
	return []string{}, false
}

func NewCache(interval time.Duration) *Cache {
	var cm = make(map[string]cacheEntry)
	cache := &Cache{
		cacheMemory: cm,
		mu:          sync.Mutex{},
	}
	go cache.readLoop(interval)
	PokedexCache = cache
	return cache
}

var PokedexCache *Cache

func GetPokedexCacheInstance() *Cache {
	return PokedexCache
}
