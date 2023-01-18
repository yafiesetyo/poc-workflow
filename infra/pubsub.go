package infra

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/yafiesetyo/poc-workflow/config"
)

type PubSubClient struct {
	Client *pubsub.Client
}

func initPubSub() (*pubsub.Client, error) {
	ctx := context.Background()
	return pubsub.NewClient(ctx, config.Cfg.GCP.ProjectID)
}

func NewPubSubClient() PubSubClient {
	client, err := initPubSub()
	if err != nil {
		log.Fatalf("failed to initialize pubsub, err: %v \n", err)
	}

	return PubSubClient{
		Client: client,
	}
}

func (c *PubSubClient) Publish(ctx context.Context, data []byte) error {
	t := c.Client.Topic(config.Cfg.GCP.PubsubTopicID)
	res := t.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err := res.Get(ctx)
	if err != nil {
		log.Default().Println("getting error when publish message, error", err)
		return err
	}

	log.Default().Println("message published")
	return nil
}
