package out

import (
	"context"
	"fraud-scoring/internal/domain/history"
	api "fraud-scoring/internal/infra/grpc"
)

type GrpcUserTransactionsRepository struct {
	grpc api.UserTransactionsServiceClient
}

func (gutr *GrpcUserTransactionsRepository) LastOrder(document string) (*history.LastOrder, error) {
	arg := &api.LastUserTransactionRequest{Document: document}
	res, err := gutr.grpc.GetLastUserTransaction(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &history.LastOrder{
		SellerId: res.SellerId,
		Currency: res.Currency,
		Amount:   res.Value,
	}, nil
}

func (gutr *GrpcUserTransactionsRepository) AverageTransactions(document string, month string) (*history.AveragePayment, error) {
	arg := &api.UserMonthAverageRequest{
		Document: document,
		Month:    month,
	}
	res, err := gutr.grpc.GetUserMonthAverage(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &history.AveragePayment{
		Month:  res.Month,
		Amount: res.Total,
	}, nil
}

func NewGrpcUserTransactionsRepository(grpc api.UserTransactionsServiceClient) *GrpcUserTransactionsRepository {
	return &GrpcUserTransactionsRepository{grpc: grpc}
}
