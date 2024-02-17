package application

import (
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/repositories"
	"fraud-scoring/internal/domain/scoring"
	"fraud-scoring/internal/domain/scoring/criteria"
)

type PaymentRiskScoring struct {
	utr repositories.UserTransactionsRepository
	tsc repositories.TransactionScoreCard
}

func (prs *PaymentRiskScoring) assessment(order *domain.PaymentOrder) {
	lastOrder, err := prs.utr.LastOrder(order.BuyerDocument)
	if err != nil {

	}
	avg, err := prs.utr.AverageTransactions(order.BuyerDocument, order.At.String())
	if err != nil {

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

	}

}

func NewPaymentRiskScoring(utr repositories.UserTransactionsRepository, tsc repositories.TransactionScoreCard) *PaymentRiskScoring {
	return &PaymentRiskScoring{utr: utr, tsc: tsc}
}
