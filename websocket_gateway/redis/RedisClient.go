package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

/*
	Redis客户端
*/

type RedisClient struct {
	RClient redis.Client
	Ctx     context.Context
	Timeout time.Duration
}

var redisClient *RedisClient

func GetRedisClient() RedisClient {
	return *redisClient
}

func init() {
	if redisClient == nil {
		redisClient = &RedisClient{}
		ctx := context.Background()
		client := redis.NewClient(&redis.Options{
			Addr:     "nas.huerpu.top:26379",
			Password: "redis&master001", // no password set
			DB:       1,                 // use default DB
		})
		redisClient.Ctx = ctx
		redisClient.RClient = *client
		redisClient.Timeout = 24 * time.Hour
	}
}

func (client RedisClient) Set(key, value string) bool {
	cmd := client.RClient.Set(client.Ctx, key, value, client.Timeout)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	return cmd.Val() == "OK"
}

func (client RedisClient) Get(key string) string {
	cmd := client.RClient.Get(client.Ctx, key)
	if cmd.Err() == redis.Nil { // Key 不存在情况
		return ""
	}
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	return cmd.Val()
}
