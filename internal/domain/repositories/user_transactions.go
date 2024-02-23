package repositories

import (
	"fraud-scoring/internal/domain/history"
	"time"
)

type UserTransactionsRepository interface {
	LastOrder(document string) (*history.LastOrder, error)
	AverageTransactions(document string, at time.Time) (*history.AveragePayment, error)
}
