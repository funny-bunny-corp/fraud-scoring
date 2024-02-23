package application

import (
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/application/errors"
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

func (prs *PaymentRiskScoring) Assessment(order *domain.TransactionAnalysis) error {
	prs.log.Info("start to performing scoring in transaction",
		zap.String("id", order.Payment.Id),
		zap.String("user_id", order.Participants.Buyer.Document),
		zap.String("seller_id", order.Participants.Seller.SellerId),
	)
	lastOrder, err := prs.utr.LastOrder(order.Participants.Buyer.Document)
	if err != nil {
		prs.log.Error("error to retrieve last transaction", zap.String("user_id", order.Participants.Buyer.Document))
		return errors.LastOrderNotFound{Err: err}
	}
	avg, err := prs.utr.AverageTransactions(order.Participants.Buyer.Document, order.Order.At)
	if err != nil {
		prs.log.Error("error to retrieve avg transaction", zap.String("user_id", order.Participants.Buyer.Document))
		return errors.AverageTransactionsNotFound{Err: err}
	}
	ac := &criteria.AverageValueCriteria{}
	sc := &criteria.SellerCriteria{Next: ac}
	cc := &criteria.CurrencyCriteria{Next: sc}
	vc := &criteria.ValueCriteria{Next: cc}
	ti := scoring.TransactionRiskScoreInput{
		Average:     avg,
		Last:        lastOrder,
		Transaction: order,
	}
	scores := &scoring.TransactionRiskFactors{}
	vc.Execute(ti, scores)

	scoreCard := &domain.ScoringResult{
		Score: domain.ScoreCard{
			ValueScore:        domain.ValueScoreCard{Score: scores.ValueScore.Scoring},
			SellerScore:       domain.SellerScoreCard{Score: scores.SellerScore.Scoring},
			AverageValueScore: domain.AverageValueScoreCard{Score: scores.AverageValue.Scoring},
			CurrencyScore:     domain.CurrencyScoreCard{Score: scores.CurrencyScore.Scoring},
		},
		Transaction: *order,
	}
	errSc := prs.tsc.Store(scoreCard)
	if errSc != nil {
		prs.log.Error("error to store scorecard in database", zap.String("user_id", order.Participants.Buyer.Document))
		return errSc
	}
	prs.log.Info("transaction was scored",
		zap.String("id", order.Payment.Id),
		zap.String("user_id", order.Participants.Buyer.Document),
		zap.String("seller_id", order.Participants.Seller.SellerId),
	)
	return nil
}

func NewPaymentRiskScoring(utr repositories.UserTransactionsRepository, tsc repositories.TransactionScoreCard, log *zap.Logger) *PaymentRiskScoring {
	return &PaymentRiskScoring{utr: utr, tsc: tsc, log: log}
}
