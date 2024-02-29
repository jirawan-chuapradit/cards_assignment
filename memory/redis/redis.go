package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

func Setup() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis to check if the connection is working
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to redis")
	return client
}
