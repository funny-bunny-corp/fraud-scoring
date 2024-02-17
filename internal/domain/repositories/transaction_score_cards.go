package repositories

import "fraud-scoring/internal/domain"

type TransactionScoreCard interface {
	Store(card *domain.TransactionScoreCard) error
}
