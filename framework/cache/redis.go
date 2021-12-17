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
			commons.Logger().Error("Redis连接成功", zap.Any("name", config.Name), zap.Any("addr", config.Addr))
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

func (cache *RedisCache) Set(key string, data interface{}, et CacheExpireTime) (result RedisResult) {
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

func (cache *RedisCache) Get(key string) (result RedisResult) {
	result.val, result.err = cache.client.Get(cache.Ctx, key).Result()
	if result.err == redis.Nil {
		commons.Logger().Warn("RedisCache.Get",
			zap.String("key", key),
			zap.String("err", "key does not exist"))
		result.err = nil
	} else if result.err != nil {
		commons.Logger().Error("RedisCache.Get",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	} else {
		result.exist = true
	}
	return
}

func (cache *RedisCache) SetNX(key string, data interface{}, et CacheExpireTime) (result RedisResult) {
	result.val, result.err = cache.client.SetNX(cache.Ctx, key, data, time.Duration(et)).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SetNX",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HDel(ctx context.Context, key string, fields ...string) (result RedisResult) {
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

func (cache *RedisCache) HExists(ctx context.Context, key, field string) (result RedisResult) {
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

func (cache *RedisCache) HGet(ctx context.Context, key, field string) (result RedisResult) {
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

func (cache *RedisCache) HGetAll(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.HGetAll(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HGetAll",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HIncrBy(ctx context.Context, key, field string, incr int64) (result RedisResult) {
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

func (cache *RedisCache) HIncrByFloat(ctx context.Context, key, field string, incr float64) (result RedisResult) {
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

func (cache *RedisCache) HKeys(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.HKeys(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HKeys",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HLen(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.HLen(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.HLen",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) HMGet(ctx context.Context, key string, fields ...string) (result RedisResult) {
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

func (cache *RedisCache) HSet(ctx context.Context, key string, values ...interface{}) (result RedisResult) {
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

func (cache *RedisCache) HMSet(ctx context.Context, key string, values ...interface{}) (result RedisResult) {
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

func (cache *RedisCache) HSetNX(ctx context.Context, key, field string, value interface{}) (result RedisResult) {
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

func (cache *RedisCache) BLPop(ctx context.Context, timeout time.Duration, keys ...string) (result RedisResult) {
	result.val, result.err = cache.client.BLPop(cache.Ctx, timeout, keys...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.BLPop",
			zap.Duration("timeout", timeout),
			zap.Strings("keys", keys),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) BRPop(ctx context.Context, timeout time.Duration, keys ...string) (result RedisResult) {
	result.val, result.err = cache.client.BRPop(cache.Ctx, timeout, keys...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.BRPop",
			zap.Duration("timeout", timeout),
			zap.Strings("keys", keys),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (result RedisResult) {
	result.val, result.err = cache.client.BRPopLPush(cache.Ctx, source, destination, timeout).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.BRPopLPush",
			zap.Duration("timeout", timeout),
			zap.String("source", source),
			zap.String("destination", destination),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (result RedisResult) {
	result.val, result.err = cache.client.LInsert(cache.Ctx, key, op, pivot, value).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LInsert",
			zap.String("key", key),
			zap.String("op", op),
			zap.Any("pivot", pivot),
			zap.Any("value", value),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LLen(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.LLen(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.BRPopLPush",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LPop(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.LPop(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LPop",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LPush(ctx context.Context, key string, values ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.LPush(cache.Ctx, key, values...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LPush",
			zap.String("key", key),
			zap.Any("values", values),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LPushX(ctx context.Context, key string, values ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.LPushX(cache.Ctx, key, values...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LPushX",
			zap.String("key", key),
			zap.Any("values", values),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LRange(ctx context.Context, key string, start, stop int64) (result RedisResult) {
	result.val, result.err = cache.client.LRange(cache.Ctx, key, start, stop).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LRange",
			zap.String("key", key),
			zap.Int64("start", start),
			zap.Int64("stop", stop),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LRem(ctx context.Context, key string, count int64, value interface{}) (result RedisResult) {
	result.val, result.err = cache.client.LRem(cache.Ctx, key, count, value).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LRem",
			zap.String("key", key),
			zap.Int64("count", count),
			zap.Any("value", value),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LSet(ctx context.Context, key string, index int64, value interface{}) (result RedisResult) {
	result.val, result.err = cache.client.LSet(cache.Ctx, key, index, value).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LSet",
			zap.String("key", key),
			zap.Int64("index", index),
			zap.Any("value", value),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) LTrim(ctx context.Context, key string, start, stop int64) (result RedisResult) {
	result.val, result.err = cache.client.LTrim(cache.Ctx, key, start, stop).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.LTrim",
			zap.String("key", key),
			zap.Int64("start", start),
			zap.Int64("stop", stop),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) RPop(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.RPop(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.RPop",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) RPopLPush(ctx context.Context, source, destination string) (result RedisResult) {
	result.val, result.err = cache.client.RPopLPush(cache.Ctx, source, destination).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.RPopLPush",
			zap.String("source", source),
			zap.String("destination", destination),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) RPush(ctx context.Context, key string, values ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.RPush(cache.Ctx, key, values...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.RPush",
			zap.String("key", key),
			zap.Any("values", values),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SAdd(ctx context.Context, key string, members ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.SAdd(cache.Ctx, key, members...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SAdd",
			zap.String("key", key),
			zap.Any("members", members),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SCard(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.SCard(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SCard",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SDiff(ctx context.Context, keys ...string) (result RedisResult) {
	result.val, result.err = cache.client.SDiff(cache.Ctx, keys...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SDiff",
			zap.Any("keys", keys),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SIsMember(ctx context.Context, key string, member interface{}) (result RedisResult) {
	result.val, result.err = cache.client.SIsMember(cache.Ctx, key, member).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SIsMember",
			zap.String("key", key),
			zap.Any("member", member),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SMIsMember(ctx context.Context, key string, members ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.SMIsMember(cache.Ctx, key, members...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SIsMember",
			zap.String("key", key),
			zap.Any("members", members),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SMembers(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.SMembers(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SMembers",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SMembersMap(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.SMembersMap(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SMembersMap",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SMove(ctx context.Context, source, destination string, member interface{}) (result RedisResult) {
	result.val, result.err = cache.client.SMove(cache.Ctx, source, destination, member).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SMove",
			zap.String("source", source),
			zap.String("destination", destination),
			zap.Any("member", member),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SPop(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.SPop(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SPop",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SPopN(ctx context.Context, key string, count int64) (result RedisResult) {
	result.val, result.err = cache.client.SPopN(cache.Ctx, key, count).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SPopN",
			zap.String("key", key),
			zap.Int64("count", count),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SRandMember(ctx context.Context, key string) (result RedisResult) {
	result.val, result.err = cache.client.SRandMember(cache.Ctx, key).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SRandMember",
			zap.String("key", key),
			zap.Any("err", result.err))
		return
	}
	return
}

// SRandMemberN Redis `SRANDMEMBER key count` command.
func (cache *RedisCache) SRandMemberN(ctx context.Context, key string, count int64) (result RedisResult) {
	result.val, result.err = cache.client.SRandMemberN(cache.Ctx, key, count).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SRandMemberN",
			zap.String("key", key),
			zap.Int64("count", count),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SRem(ctx context.Context, key string, members ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.SRem(cache.Ctx, key, members...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SRem",
			zap.String("key", key),
			zap.Any("members", members),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) SUnion(ctx context.Context, keys ...string) (result RedisResult) {
	result.val, result.err = cache.client.SUnion(cache.Ctx, keys...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.SUnion",
			zap.Any("keys", keys),
			zap.Any("err", result.err))
		return
	}
	return
}

func (cache *RedisCache) Do(args ...interface{}) (result RedisResult) {
	result.val, result.err = cache.client.Do(cache.Ctx, args...).Result()
	if result.err != nil {
		commons.Logger().Error("RedisCache.Do",
			zap.Any("args", args),
			zap.Any("err", result.err))
		return
	}
	return
}
