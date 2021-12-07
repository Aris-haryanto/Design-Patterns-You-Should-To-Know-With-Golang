package adapters

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisAdapter struct {
	Host     string
	Password string
}

var ctx = context.Background()

func (ra *RedisAdapter) Connect() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     ra.Host,
		Password: ra.Password,
	})

	return redisClient
}

func (ra *RedisAdapter) Publish(channel string, message string) {
	c := ra.Connect()

	pubsub := c.Subscribe(ctx, channel)

	c.Publish(ctx, channel, message)

	msg, errRcv := pubsub.ReceiveMessage(ctx)
	if errRcv != nil {
		panic(errRcv)
	}

	fmt.Println(msg.Payload)
}
