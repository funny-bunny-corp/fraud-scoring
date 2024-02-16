package scoring

type Rule interface {
	Execute(input TransactionRiskScoreInput, factors *TransactionRiskFactors)
}
