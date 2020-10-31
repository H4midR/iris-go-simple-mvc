package myredis

import (
	"github.com/go-redis/redis"
)

func NewRedisCliet() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

	// Output: PONG <nil>
}
