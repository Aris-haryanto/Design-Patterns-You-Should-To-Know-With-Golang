package adapters

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type NatsAdapters struct {
	Host string
}

func (na *NatsAdapters) Connect() *nats.Conn {
	conn, err := nats.Connect(na.Host)
	if err != nil {
		panic(err)
	}

	return conn
}

func (na *NatsAdapters) Publish(channel string, message string) {
	// call connection
	c := na.Connect()

	// set subscriptions
	c.Subscribe(channel, func(m *nats.Msg) {
		fmt.Println(string(m.Data))
	})

	//publish to service receiver
	c.Publish(channel, []byte(message))
}
