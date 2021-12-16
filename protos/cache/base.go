package cache

import "fmt"

type CacheKeys struct {
	prefix string
}

func (keys *CacheKeys) format(key string, values ...interface{}) string {
	key = fmt.Sprintf("%s:%s", keys.prefix, key)
	return fmt.Sprintf(key, values)
}
