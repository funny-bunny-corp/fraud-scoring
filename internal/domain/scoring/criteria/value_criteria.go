package criteria

import "fraud-scoring/internal/domain/scoring"

type ValueCriteria struct {
	Next scoring.Rule
}

func (v *ValueCriteria) Execute(input scoring.TransactionRiskScoreInput, factors *scoring.TransactionRiskFactors) {
	if input.Order.Amount == input.Last.Amount {
		factors.WithValueScore(scoring.ValueRiskScoreEvaluation{Scoring: -3})
	} else {
		factors.WithValueScore(scoring.ValueRiskScoreEvaluation{Scoring: 0})
	}
	if v.Next != nil {
		v.Next.Execute(input, factors)
	}
}
