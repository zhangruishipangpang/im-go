package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
	"time"
)

func TestRedisClient(t *testing.T) {
	init1()
}

func init1() {
	redisClient := GetRedisClient()
	fmt.Println(redisClient)

	cmd := redisClient.Set("AK", "AV")
	fmt.Println("set : ", cmd)

	get := redisClient.Get("AK")
	fmt.Println("get:", get)
}

func client() {

	ctx := context.Background()
	var client *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "nas.huerpu.top:26379",
		Password: "redis&master001", // no password set
		DB:       1,                 // use default DB
	})
	cmd := client.Set(ctx, "TK2", "TV", 900*time.Hour)
	fmt.Println(cmd.Result())

	get := client.Get(ctx, "TK")
	fmt.Println("获取的值； ", get.Val())

}
