package services

type IAdapter interface {
	Publish(channel string, message string)
}

type PubSub struct {
	Adapter IAdapter
}

// function ini digunakan untuk set data ke struct yang ada di function interface
func (pb *PubSub) SetStruct(data IAdapter) {
	pb.Adapter = data
}

// function ini digunakan untuk binding fungsi publish ke function sebenarnya di folder /adapters sesuai dengan adapter nya
func (pb *PubSub) Publish(channel string, message string) {
	pb.Adapter.Publish(channel, message)
}
