package adapters

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisAdapter struct {
	Conn *redis.Client
}

func RedisConn(host string, password string) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})

	return redisClient
}

func (ra *RedisAdapter) Publish(channel string, message string) {
	// set context timeout if publish more than 1 minute
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	//publish message with go routine for asynch
	ra.Conn.Publish(ctx, channel, message)
}

func (ra *RedisAdapter) Listener(channel string) {
	// set context cancel if listener goes wrong
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pubsub := ra.Conn.Subscribe(ctx, channel)

	for {
		msg, errRcv := pubsub.ReceiveMessage(ctx)
		if errRcv != nil {
			log.Println(errRcv)
		}

		log.Println(msg.Payload)
	}
}
