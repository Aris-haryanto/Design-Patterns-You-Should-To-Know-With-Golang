package main

import (
	"context"
	"log"

	"builder-creational-patterns/services"
)

func main() {
	ctx := context.Background()

	// this is builder pattern start
	// we create sepearate function and call them separately to build some service
	pubSubSvc := services.GooglePubsubListener{
		Ctx:              ctx,
		ProjectID:        "google project",         // google project ID
		SubsID:           "topic pubsub subscribe", // pubsub topic subscribe
		PubSubCredential: "./credential.json",      // credential from google
	}

	pubSubSvc.New()
	pubSubSvc.Subscribe()

	if err := pubSubSvc.Run(); err != nil {
		log.Println(err)
	}

	defer pubSubSvc.Close()

}
