package scoring

type TransactionRiskFactors struct {
	SellerScore   SellerRiskScoreEvaluation
	CurrencyScore CurrencyRiskScoreEvaluation
	ValueScore    ValueRiskScoreEvaluation
	AverageValue  AverageValueRiskScoreEvaluation
}

type SellerRiskScoreEvaluation struct {
	Scoring int
}

type CurrencyRiskScoreEvaluation struct {
	Scoring int
}

type ValueRiskScoreEvaluation struct {
	Scoring int
}

type AverageValueRiskScoreEvaluation struct {
	Scoring int
}

func (trf TransactionRiskFactors) WithCurrencyScore(crse CurrencyRiskScoreEvaluation) {
	trf.CurrencyScore = crse
}

func (trf TransactionRiskFactors) WithSellerScore(srse SellerRiskScoreEvaluation) {
	trf.SellerScore = srse
}
