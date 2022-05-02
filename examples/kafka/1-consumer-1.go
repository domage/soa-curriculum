package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, "key").Result()

	for err == redis.Nil {
		val, err = rdb.Get(ctx, "key").Result()

		time.Sleep(2 * time.Second)
	}

	fmt.Println("Value from DB", val)
	rdb.Del(ctx, "key")
}
