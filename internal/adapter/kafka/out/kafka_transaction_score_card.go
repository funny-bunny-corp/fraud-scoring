package out

import (
	"context"
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/infra/kafka"
	"github.com/IBM/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"log"
)

const (
	eventType    = "paymentic.io.fraud-detection.v1.transaction.scorecard.created"
	eventSource  = "fraud-scoring"
	eventSubject = "score-card-ready"
)

type KafkaTransactionScoreCard struct {
	cli kafka.CloudEventsSender
}

func (ktsc *KafkaTransactionScoreCard) Store(card *domain.TransactionScoreCard) error {
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetType(eventType)
	e.SetSource(eventSource)
	e.SetSubject(eventSubject)
	_ = e.SetData(cloudevents.ApplicationJSON, card)
	if result := ktsc.cli.Send(
		kafka_sarama.WithMessageKey(context.Background(), sarama.StringEncoder(e.ID())),
		e,
	); cloudevents.IsUndelivered(result) {
		log.Printf("failed to send: %v", result)
	} else {
		log.Printf("sent: %s, accepted: %t", e.ID(), cloudevents.IsACK(result))
	}
	return nil
}

func NewKafkaTransactionScoreCard(cli kafka.CloudEventsSender) *KafkaTransactionScoreCard {
	return &KafkaTransactionScoreCard{cli: cli}
}
