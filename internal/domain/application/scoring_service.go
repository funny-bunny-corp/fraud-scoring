package application

import (
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/errors"
	"fraud-scoring/internal/domain/repositories"
	"fraud-scoring/internal/domain/scoring"
	"fraud-scoring/internal/domain/scoring/criteria"
	"go.uber.org/zap"
)

type PaymentRiskScoring struct {
	utr repositories.UserTransactionsRepository
	tsc repositories.TransactionScoreCard
	log *zap.Logger
}

func (prs *PaymentRiskScoring) Assessment(order *domain.PaymentOrder) error {
	prs.log.Info("start to performing scoring in transaction",
		zap.String("id", order.Id),
		zap.String("user_id", order.BuyerDocument),
		zap.String("seller_id", order.SellerId),
	)
	lastOrder, err := prs.utr.LastOrder(order.BuyerDocument)
	if err != nil {
		prs.log.Error("error to retrieve last transaction", zap.String("user_id", order.BuyerDocument))
		return errors.LastOrderNotFound{Err: err}
	}
	avg, err := prs.utr.AverageTransactions(order.BuyerDocument, order.At.String())
	if err != nil {
		prs.log.Error("error to retrieve avg transaction", zap.String("user_id", order.BuyerDocument))
		return errors.AverageTransactionsNotFound{Err: err}
	}
	ac := &criteria.AverageValueCriteria{}
	sc := &criteria.SellerCriteria{Next: ac}
	cc := &criteria.CurrencyCriteria{Next: sc}
	vc := &criteria.ValueCriteria{Next: cc}
	ti := scoring.TransactionRiskScoreInput{
		Average: avg,
		Last:    lastOrder,
		Order:   order,
	}
	scores := &scoring.TransactionRiskFactors{}
	vc.Execute(ti, scores)
	scoreCard := &domain.TransactionScoreCard{
		Transaction:       domain.Transaction{Id: order.Id},
		ValueScore:        domain.ValueScoreCard{Score: scores.ValueScore.Scoring},
		SellerScore:       domain.SellerScoreCard{Score: scores.SellerScore.Scoring},
		AverageValueScore: domain.AverageValueScoreCard{Score: scores.AverageValue.Scoring},
		CurrencyScore:     domain.CurrencyScoreCard{Score: scores.CurrencyScore.Scoring},
	}
	errSc := prs.tsc.Store(scoreCard)
	if errSc != nil {
		prs.log.Error("error to store scorecard in database", zap.String("user_id", order.BuyerDocument))
		return errSc
	}
	prs.log.Info("transaction was scored",
		zap.String("id", order.Id),
		zap.String("user_id", order.BuyerDocument),
		zap.String("seller_id", order.SellerId),
	)
	return nil
}

func NewPaymentRiskScoring(utr repositories.UserTransactionsRepository, tsc repositories.TransactionScoreCard, log *zap.Logger) *PaymentRiskScoring {
	return &PaymentRiskScoring{utr: utr, tsc: tsc, log: log}
}
