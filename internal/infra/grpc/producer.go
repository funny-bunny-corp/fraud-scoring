package api

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type UserTransactionsConfig struct {
	Host string
}

func NewUserTransactionGrpc(config *UserTransactionsConfig) UserTransactionsServiceClient {
	conn, err := grpc.Dial(config.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

	}
	client := NewUserTransactionsServiceClient(conn)
	return client
}

func NewUserTransactionsConfig() *UserTransactionsConfig {
	return &UserTransactionsConfig{Host: os.Getenv("USER_TRANSACTIONS_HOST")}
}
