package external

import (
	"HashChallenge/logger"
	"sync"
)

type InMemoryCache struct {
	cache map[string]interface{}
	mutex sync.RWMutex
}

var once sync.Once
var (
	instance InMemoryCache
)

func GetCacheInstance() *InMemoryCache {
	once.Do(func() { // atomic, does not allow repeating
		instance = InMemoryCache{}
		logger.PushLogs("cache initialiazed", logger.Level.Info)
		instance.cache = make(map[string]interface{}) // thread safe
	})
	return &instance
}
func (c *InMemoryCache) get(key string) (interface{}, bool) {
	value, err := c.cache[key]
	return value, err
}
func (c *InMemoryCache) put(key string, value interface{}) {
	c.mutex.Lock()
	if _, ok := c.cache[key]; !ok {
		c.cache[key] = value
	}
	c.mutex.Unlock()
}
