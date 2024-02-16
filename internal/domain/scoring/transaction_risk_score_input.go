package scoring

import (
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/history"
)

type TransactionRiskScoreInput struct {
	Average history.AveragePayment
	Last    history.LastOrder
	Order   domain.PaymentOrder
}
