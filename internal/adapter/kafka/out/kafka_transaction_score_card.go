package out

import (
	"context"
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/infra/kafka"
	"github.com/IBM/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	eventType    = "paymentic.io.fraud-detection.v1.transaction.scorecard.created"
	eventSource  = "fraud-scoring"
	eventSubject = "score-card-ready"
)

type KafkaTransactionScoreCard struct {
	cli kafka.CloudEventsSender
	log *zap.Logger
}

func (ktsc *KafkaTransactionScoreCard) Store(card *domain.ScoringResult) error {
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
		ktsc.log.Error("failed to send", zap.String("error", result.Error()))
	} else {
		ktsc.log.Info("message sent", zap.String("id", e.ID()), zap.Bool("ack", cloudevents.IsACK(result)))
	}
	return nil
}

func NewKafkaTransactionScoreCard(cli kafka.CloudEventsSender, log *zap.Logger) *KafkaTransactionScoreCard {
	return &KafkaTransactionScoreCard{cli: cli, log: log}
}
