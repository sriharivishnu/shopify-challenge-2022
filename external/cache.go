package external

import "time"

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{}, ttl int64)
}

type entry struct {
	Value     interface{}
	TTL       int64
	CreatedAt int64
}

// Could implement external cache using redis, memcached, etc.
type InMemoryCache struct {
	InMemoryDs map[string]entry
}

func (cache *InMemoryCache) Init() {
	cache.InMemoryDs = make(map[string]entry)
}

func (cache *InMemoryCache) Get(key string) interface{} {
	if entry, ok := cache.InMemoryDs[key]; ok {
		if entry.TTL > 0 && entry.TTL < time.Now().Unix() {
			delete(cache.InMemoryDs, key)
			return nil
		}
		return entry.Value
	}
	return nil
}

func (cache *InMemoryCache) Set(key string, value interface{}, ttl int64) {
	cache.InMemoryDs[key] = entry{Value: value, TTL: ttl, CreatedAt: time.Now().Unix()}
}
