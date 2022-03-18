package adapters

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatsAdapter struct {
	Conn *nats.Conn
}

func NatsConn(host string) *nats.Conn {
	conn, err := nats.Connect(host)
	if err != nil {
		log.Println(err)
	}

	return conn
}

func (na *NatsAdapter) Publish(channel string, message string) {
	//publish to service receiver
	na.Conn.Publish(channel, []byte(message))
}

func (na *NatsAdapter) Listener(channel string) {
	// set subscriptions
	na.Conn.Subscribe(channel, func(m *nats.Msg) {
		log.Println(string(m.Data))
	})
}
