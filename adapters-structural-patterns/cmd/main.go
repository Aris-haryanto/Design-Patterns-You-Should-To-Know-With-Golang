package main

import (
	"adapters-structural-patterns/adapters"
	"adapters-structural-patterns/services"
)

const (
	channel = "test_channel"
)

//set data nats connection ke struct Pubsub yang ada di /services
func NatsConn() *services.PubSub {
	nConn := adapters.NatsConn("127.0.0.1:4222")
	setConn := &adapters.NatsAdapter{Conn: nConn}
	return &services.PubSub{
		Adapter: setConn,
	}
}

//set data redis connection ke struct Pubsub yang ada di /services
func RedisConn() *services.PubSub {
	rConn := adapters.RedisConn("127.0.0.1:6379", "")
	setConn := &adapters.RedisAdapter{Conn: rConn}
	return &services.PubSub{
		Adapter: setConn,
	}
}

func main() {

	//define struct pubsub dari package /service
	pb := services.PubSub{}

	//set menggunakan adapter nats
	pb.SetStruct(NatsConn())
	go pb.Publish(channel, "ini publish dari nats")
	pb.Listener(channel)

	//====================

	//set menggunakan adapter redis
	pb.SetStruct(RedisConn())
	go pb.Publish(channel, "ini publish dari redis")
	pb.Listener(channel)
}
