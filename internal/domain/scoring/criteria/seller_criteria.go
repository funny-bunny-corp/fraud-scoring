package criteria

import "fraud-scoring/internal/domain/scoring"

type SellerCriteria struct {
	Next scoring.Rule
}

func (s *SellerCriteria) Execute(input scoring.TransactionRiskScoreInput, factors *scoring.TransactionRiskFactors) {
	if input.Order.SellerId == input.Last.SellerId {
		factors.WithSellerScore(scoring.SellerRiskScoreEvaluation{Scoring: -1})
	} else {
		factors.WithSellerScore(scoring.SellerRiskScoreEvaluation{Scoring: 0})
	}
	if s.Next != nil {
		s.Next.Execute(input, factors)
	}
}
