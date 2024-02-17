package repositories

import "fraud-scoring/internal/domain/history"

type UserTransactionsRepository interface {
	LastOrder(document string) (*history.LastOrder, error)
	AverageTransactions(document string, month string) (*history.AveragePayment, error)
}
