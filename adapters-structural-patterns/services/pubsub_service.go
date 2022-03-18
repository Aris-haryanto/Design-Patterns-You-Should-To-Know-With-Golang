package services

import "time"

type IAdapter interface {
	Publish(channel string, message string)
	Listener(channel string)
	//add another function in here
}

type PubSub struct {
	Adapter IAdapter
}

// function ini digunakan untuk set data ke struct yang ada di function interface
func (pb *PubSub) SetAdapter(data IAdapter) {
	pb.Adapter = data
}

// function ini digunakan untuk binding fungsi publish ke function sebenarnya di folder /adapters sesuai dengan adapter nya
func (pb *PubSub) Publish(channel string, message string) {
	//kita kasih jeda 1 detik sebelum publish agar listenernya ready dulu
	time.Sleep(1 * time.Second)
	pb.Adapter.Publish(channel, message)
}

func (pb *PubSub) Listener(channel string) {
	pb.Adapter.Listener(channel)
}
