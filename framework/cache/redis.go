package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"goku.net/framework/commons"
	"goku.net/framework/config"
)

type RedisClient struct {
	Ctx    context.Context
	client *redis.Client
}

func (client *RedisClient) GetClient() *redis.Client {
	return client.client
}

var (
	redisClients map[string]*RedisClient
)

func Init(configs []*config.RedisConfig) {
	redisClients = make(map[string]*RedisClient)
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
			redisClients[config.Name] = &RedisClient{
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

func GetClient(name string) *RedisClient {
	if client, ok := redisClients[name]; ok {
		return client
	}
	return nil
}
