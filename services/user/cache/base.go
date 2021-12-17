package cache

import (
	"goku.net/framework/cache"
)

func Cache() *cache.RedisCache {
	return cache.Cache("user")
}
