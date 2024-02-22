package criteria

import "fraud-scoring/internal/domain/scoring"

type AverageValueCriteria struct {
	Next scoring.Rule
}

func (a *AverageValueCriteria) Execute(input scoring.TransactionRiskScoreInput, factors *scoring.TransactionRiskFactors) {
	if input.Transaction.Payment.Amount >= input.Average.Amount {
		factors.WithAverageValueScore(scoring.AverageValueRiskScoreEvaluation{Scoring: -3})
	} else {
		factors.WithAverageValueScore(scoring.AverageValueRiskScoreEvaluation{Scoring: 0})
	}
	if a.Next != nil {
		a.Next.Execute(input, factors)
	}
}
