package criteria

import "fraud-scoring/internal/domain/scoring"

type CurrencyCriteria struct {
	Next scoring.Rule
}

func (c *CurrencyCriteria) Execute(input scoring.TransactionRiskScoreInput, factors *scoring.TransactionRiskFactors) {
	if input.Order.Currency != input.Last.Currency {
		factors.WithCurrencyScore(scoring.CurrencyRiskScoreEvaluation{Scoring: -1})
	} else {
		factors.WithCurrencyScore(scoring.CurrencyRiskScoreEvaluation{Scoring: 0})
	}
	if c.Next != nil {
		c.Next.Execute(input, factors)
	}
}
