package main

import (
	"adapters-structural-patterns/adapters"
	"adapters-structural-patterns/services"
	"runtime"
)

const (
	channel = "test_channel"
)

//set data nats connection ke struct Pubsub yang ada di /services
func NatsConn() *services.PubSub {
	nconn := &adapters.NatsAdapters{Host: "127.0.0.1"}
	return &services.PubSub{
		Adapter: nconn,
	}
}

//set data redis connection ke struct Pubsub yang ada di /services
func RedisConn() *services.PubSub {
	nconn := &adapters.RedisAdapter{Host: "127.0.0.1:6379", Password: ""}
	return &services.PubSub{
		Adapter: nconn,
	}
}

func main() {

	//define struct pubsub dari package /service
	nat := services.PubSub{}

	//publish from nats
	nat.SetStruct(NatsConn())
	nat.Publish(channel, "ini publish dari nats")

	//publish from redis
	nat.SetStruct(RedisConn())
	nat.Publish(channel, "ini publish dari redis")

	// Keep the connection alive
	runtime.Goexit()
}
