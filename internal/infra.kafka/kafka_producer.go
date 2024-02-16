package infra_kafka

import (
	"github.com/IBM/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"

	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
)

func NewCloudEventsKafkaConsumer(sc *SaramaConfig) (cloudevents.Client, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	receiver, err := kafka_sarama.NewConsumer([]string{sc.Host}, saramaConfig, sc.GroupId, sc.Topic)
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}
	c, err := cloudevents.NewClient(receiver)
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	return c, nil
}
