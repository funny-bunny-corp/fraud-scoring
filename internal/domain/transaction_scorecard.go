package domain

type ScoringResult struct {
	Score       ScoreCard           `json:"score"`
	Transaction TransactionAnalysis `json:"transaction"`
}

type ScoreCard struct {
	ValueScore        ValueScoreCard        `json:"value_score"`
	SellerScore       SellerScoreCard       `json:"seller_score"`
	AverageValueScore AverageValueScoreCard `json:"average_value_score"`
	CurrencyScore     CurrencyScoreCard     `json:"currency_score"`
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
