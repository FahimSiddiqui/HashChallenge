package external

import (
	"testing"
)

func TestCacheOperations(t *testing.T) {
	_, ok := GetCacheInstance().get("key1")
	// Since there are not values in the cache, it must miss.
	if ok {
		t.Error()
	}

	// Now lets put the value in cache
	GetCacheInstance().put("key1", "value1")

	// Now, cache must hit and return value for Key1
	value, ok := GetCacheInstance().get("key1")
	if value != "value1" {
		t.Fail()
	}
}
