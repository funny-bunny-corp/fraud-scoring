package domain

type ScoringResult struct {
	Score       ScoreCard           `json:"score"`
	Transaction TransactionAnalysis `json:"transaction"`
}

type ScoreCard struct {
	ValueScore        ValueScoreCard        `json:"valueScore"`
	SellerScore       SellerScoreCard       `json:"sellerScore"`
	AverageValueScore AverageValueScoreCard `json:"averageValueScore"`
	CurrencyScore     CurrencyScoreCard     `json:"currencyScore"`
}

type Transaction struct {
	Id string `json:"id"`
}

type ValueScoreCard struct {
	Score int `json:"score"`
}

type SellerScoreCard struct {
	Score int `json:"score"`
}

type AverageValueScoreCard struct {
	Score int `json:"score"`
}

type CurrencyScoreCard struct {
	Score int `json:"score"`
}
