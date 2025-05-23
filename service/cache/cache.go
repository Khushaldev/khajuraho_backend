package cache

import (
	"khajuraho/backend/utils"
	"time"
)

const (
	CategoriesKey = "categories"
)

type Cache interface {
	Set(key, value string, expiration time.Duration) error
	Get(key string) (string, bool)
}

var cacheInstance Cache

func InitCache() {
	cacheInstance = NewInMemoryCache()
}

func Get(key string) (string, bool) {
	return cacheInstance.Get(key)
}

func Set(key string, value string) error {
	return cacheInstance.Set(key, value, 20*time.Second)
}

// GetOrSetCache fetches from cache or uses fallback and stores in cache.
// Ensures fallback return type matches *target.
func GetOrSetCache[T any](key string, target *T, fallback func() (T, error)) error {
	value, found := Get(key)
	if found {
		return utils.Unmarshal(value, target)
	}

	fallBackData, err := fallback()
	if err != nil {
		return err
	}

	*target = fallBackData
	strVal, err := utils.Marshal(fallBackData)
	if err == nil {
		Set(key, strVal)
	}

	return nil
}
