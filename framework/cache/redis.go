package cache

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"goku.net/framework/commons"
	"goku.net/framework/config"
)

type RedisCache struct {
	Ctx    context.Context
	client *redis.Client
}

var (
	redisClients map[string]*RedisCache
)

func Init(configs []*config.RedisConfig) {
	redisClients = make(map[string]*RedisCache)
	for _, config := range configs {
		if len(config.Name) == 0 {
			config.Name = "default"
		}
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.Db,
		})
		if client != nil {
			redisClients[config.Name] = &RedisCache{
				Ctx:    context.Background(),
				client: client,
			}
		} else {
			commons.Logger().Error("Redis连接异常", zap.Any("config", config))
			os.Exit(-1)
			return
		}
	}
}

func DefaultCache() *RedisCache {
	if client, ok := redisClients["default"]; ok {
		return client
	}
	return nil
}

func Cache(name string) *RedisCache {
	if client, ok := redisClients[name]; ok {
		return client
	}
	return nil
}

func (cache *RedisCache) GetClient() *redis.Client {
	return cache.client
}

func (cache *RedisCache) Set(key string, data interface{}, et CacheExpireTime) (result *RedisResult) {
	result.err = cache.client.Set(cache.Ctx, key, data, time.Duration(et)).Err()
	if result.err != nil {
		commons.Logger().Error("RedisCache.Set",
			zap.String("key", key),
			zap.Any("data", data),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) Get(key string) (result *RedisResult) {
	result.val, result.err = cache.client.Get(cache.Ctx, key).Result()
	if result.err == redis.Nil {
		commons.Logger().Warn("RedisCache.Get",
			zap.String("key", key),
			zap.String("err", "key does not exist"))
	} else if result.err != nil {
		commons.Logger().Error("RedisCache.Get",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SetNX(key string, data interface{}, et CacheExpireTime) (result *RedisResult) {
	result.val, result.err = cache.client.SetNX(cache.Ctx, key, data, time.Duration(et)).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SetNX",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HDel(ctx context.Context, key string, fields ...string) (result *RedisResult) {
	result.val, result.err = cache.client.HDel(cache.Ctx, key, fields...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HDel",
			zap.String("key", key),
			zap.Strings("fields", fields),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HExists(ctx context.Context, key, field string) (result *RedisResult) {
	result.val, result.err = cache.client.HExists(cache.Ctx, key, field).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HExists",
			zap.String("key", key),
			zap.String("field", field),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HGet(ctx context.Context, key, field string) (result *RedisResult) {
	result.val, result.err = cache.client.HGet(cache.Ctx, key, field).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HGet",
			zap.String("key", key),
			zap.String("field", field),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HGetAll(ctx context.Context, key string) (result *RedisResult) {
	result.val, result.err = cache.client.HGetAll(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HGetAll",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HIncrBy(ctx context.Context, key, field string, incr int64) (result *RedisResult) {
	result.val, result.err = cache.client.HIncrBy(cache.Ctx, key, field, incr).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HIncrBy",
			zap.String("key", key),
			zap.String("field", field),
			zap.Int64("incr", incr),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HIncrByFloat(ctx context.Context, key, field string, incr float64) (result *RedisResult) {
	result.val, result.err = cache.client.HIncrByFloat(cache.Ctx, key, field, incr).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HIncrByFloat",
			zap.String("key", key),
			zap.String("field", field),
			zap.Float64("incr", incr),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HKeys(ctx context.Context, key string) (result *RedisResult) {
	result.val, result.err = cache.client.HKeys(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HKeys",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HLen(ctx context.Context, key string) (result *RedisResult) {
	result.val, result.err = cache.client.HLen(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HLen",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HMGet(ctx context.Context, key string, fields ...string) (result *RedisResult) {
	result.val, result.err = cache.client.HMGet(cache.Ctx, key, fields...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HMGet",
			zap.String("key", key),
			zap.Strings("fields", fields),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HSet(ctx context.Context, key string, values ...interface{}) (result *RedisResult) {
	result.val, result.err = cache.client.HSet(cache.Ctx, key, values...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HSet",
			zap.String("key", key),
			zap.Any("values", values),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HMSet(ctx context.Context, key string, values ...interface{}) (result *RedisResult) {
	result.val, result.err = cache.client.HMSet(cache.Ctx, key, values...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HMSet",
			zap.String("key", key),
			zap.Any("values", values),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HSetNX(ctx context.Context, key, field string, value interface{}) (result *RedisResult) {
	result.val, result.err = cache.client.HSetNX(cache.Ctx, key, field, value).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HSetNX",
			zap.String("key", key),
			zap.String("field", field),
			zap.Any("value", value),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) Do(args ...interface{}) (result *RedisResult) {
	result.val, result.err = cache.client.Do(cache.Ctx, args...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.Do",
			zap.Any("args", args),
			zap.Any("err", result.err))
		return
	}
	return
}
