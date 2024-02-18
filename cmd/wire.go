//go:build wireinject
// +build wireinject

package main

import (
	out2 "fraud-scoring/internal/adapter/grpc/out"
	"fraud-scoring/internal/adapter/kafka/in"
	"fraud-scoring/internal/adapter/kafka/out"
	"fraud-scoring/internal/domain/application"
	"fraud-scoring/internal/domain/repositories"
	api "fraud-scoring/internal/infra/grpc"
	ik "fraud-scoring/internal/infra/kafka"
	"fraud-scoring/internal/infra/logger"
	"github.com/google/wire"
)

func buildAppContainer() (*Manager, error) {
	wire.Build(ik.NewSaramaConfig,
		api.NewUserTransactionsConfig,
		api.NewUserTransactionGrpc,
		ik.NewCloudEventsKafkaSender,
		ik.NewCloudEventsKafkaConsumer,
		out.NewKafkaTransactionScoreCard,
		logger.NewLogger,
		out2.NewGrpcUserTransactionsRepository,
		wire.Bind(new(repositories.TransactionScoreCard), new(*out.KafkaTransactionScoreCard)),
		wire.Bind(new(repositories.UserTransactionsRepository), new(*out2.GrpcUserTransactionsRepository)),
		application.NewPaymentRiskScoring,
		in.NewCheckoutEventReceiver,
		NewManager,
	)
	return nil, nil
}
