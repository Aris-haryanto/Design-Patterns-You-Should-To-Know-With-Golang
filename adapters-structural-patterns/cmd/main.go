package main

import (
	"adapters-structural-patterns/adapters"
	"adapters-structural-patterns/services"
	"runtime"
)

const (
	channel = "test_channel"
	// message = "test message"
)

func NatsConn() *services.PubSub {
	nconn := &adapters.NatsAdapters{Host: "127.0.0.1"}
	return &services.PubSub{
		Adapter: nconn,
	}
}

func RedisConn() *services.PubSub {
	nconn := &adapters.RedisAdapter{Host: "127.0.0.1:6379", Password: ""}
	return &services.PubSub{
		Adapter: nconn,
	}
}

func main() {

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
