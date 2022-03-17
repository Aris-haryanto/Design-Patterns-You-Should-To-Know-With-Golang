package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type GooglePubsubListener struct {
	client           *pubsub.Client
	subscription     *pubsub.Subscription
	Ctx              context.Context
	ProjectID        string
	SubsID           string
	PubSubCredential string
}

type Msg struct {
	//this struct must same with data sent from pub/sub
	Username string `json:"username,omitempty"`
}

func (listener *GooglePubsubListener) New() {
	conn, err := pubsub.NewClient(listener.Ctx, listener.ProjectID, option.WithCredentialsFile(listener.PubSubCredential))
	if err != nil {
		log.Println(err.Error())
	}

	//set pubsub connection to struct
	listener.client = conn
}

func (listener *GooglePubsubListener) Subscribe() {
	sub := listener.client.Subscription(listener.SubsID)

	// you can set another config here
	sub.ReceiveSettings.MaxOutstandingMessages = 10

	listener.subscription = sub

}

func (listener *GooglePubsubListener) Run() error {

	log.Println("Starting Pubsub listener ...")

	ctx, cancel := context.WithCancel(listener.Ctx)
	err := listener.subscription.Receive(ctx, func(c context.Context, m *pubsub.Message) {

		var msg Msg
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println("Hi user " + msg.Username)

		// ack data when finish proccess
		m.Ack()
	})

	if err != nil {
		// print error
		log.Println(err.Error())

		//cancel process
		cancel()

		return err
	}

	return nil
}

func (listener *GooglePubsubListener) Close() {
	log.Println("Closing Pubsub Connection")
	listener.client.Close()
}
