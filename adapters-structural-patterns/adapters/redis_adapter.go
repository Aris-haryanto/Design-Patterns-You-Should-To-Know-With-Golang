package adapters

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisAdapter struct {
	Host     string
	Password string
}

func (ra *RedisAdapter) Connect() (redis.Conn, error) {
	const healthCheckPeriod = time.Minute
	return redis.Dial("tcp", ra.Host,
		// Read timeout on server should be greater than ping period.
		redis.DialReadTimeout(healthCheckPeriod+10*time.Second),
		redis.DialWriteTimeout(10*time.Second),
		redis.DialPassword(ra.Password))
}

func (ra *RedisAdapter) Publish(channel string, content string) {
	c, err := ra.Connect()
	if err != nil {
		panic(err)
	}

	pb := redis.PubSubConn{c}

	// Set up subscriptions
	pb.Subscribe(channel)

	// While not a permanent error on the connection.
	for c.Err() == nil {
		switch v := pb.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Printf("gagal")
		}
	}

	c.Do("PUBLISH", channel, content)
}
