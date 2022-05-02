package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pubsub := rdb.PSubscribe(ctx, "channel_*")
	defer pubsub.Close()

	msgi, _ := pubsub.Receive(ctx)
	subscr := msgi.(*redis.Subscription)

	fmt.Println("Received: ", subscr.String())

	msgi, _ = pubsub.Receive(ctx)
	msg := msgi.(*redis.Message)

	fmt.Println("Received: ", msg.Payload, msg.Channel)
}
