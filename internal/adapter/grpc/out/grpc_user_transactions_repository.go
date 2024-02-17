package out

import (
	"fraud-scoring/internal/domain/history"
	api "fraud-scoring/internal/infra/grpc"
)

type GrpcUserTransactionsRepository struct {
	grpc *api.UserTransactionsServiceClient
}

func (gutr *GrpcUserTransactionsRepository) LastOrder(document string) (*history.LastOrder, error) {
	return nil, nil
}

func (gutr *GrpcUserTransactionsRepository) AverageTransactions(document string, month string) (*history.AveragePayment, error) {
	return nil, nil
}

func NewGrpcUserTransactionsRepository(grpc *api.UserTransactionsServiceClient) *GrpcUserTransactionsRepository {
	return &GrpcUserTransactionsRepository{grpc: grpc}
}
