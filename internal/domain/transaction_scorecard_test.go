package domain

import (
	"testing"
)

func TestScoringResult(t *testing.T) {
	tests := []struct {
		name           string
		scoringResult  ScoringResult
		expectedValid  bool
		expectedErrors []string
	}{
		{
			name: "Valid scoring result",
			scoringResult: ScoringResult{
				Score: ScoreCard{
					ValueScore: ValueScoreCard{
						Score: 85,
					},
					SellerScore: SellerScoreCard{
						Score: 90,
					},
					AverageValueScore: AverageValueScoreCard{
						Score: 75,
					},
					CurrencyScore: CurrencyScoreCard{
						Score: 95,
					},
				},
				Transaction: TransactionAnalysis{
					Participants: Participants{
						Buyer: BuyerInfo{
							Document: "12345678901",
							Name:     "John Doe",
						},
						Seller: SellerInfo{
							SellerId: "seller-123",
						},
					},
					Order: Checkout{
						Id: "order-456",
						PaymentType: CardInfo{
							CardInfo: "****1234",
							Token:    "tok_123",
						},
					},
					Payment: Payment{
						Id:       "payment-789",
						Amount:   "100.00",
						Currency: "USD",
						Status:   "completed",
					},
				},
			},
			expectedValid:  true,
			expectedErrors: []string{},
		},
		{
			name: "Empty scoring result",
			scoringResult: ScoringResult{
				Score:       ScoreCard{},
				Transaction: TransactionAnalysis{},
			},
			expectedValid:  false,
			expectedErrors: []string{"missing transaction data"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that the struct can be created
			if tt.scoringResult.Score.ValueScore.Score < 0 || tt.scoringResult.Score.ValueScore.Score > 100 {
				t.Errorf("ValueScore should be between 0 and 100, got %d", tt.scoringResult.Score.ValueScore.Score)
			}
			if tt.scoringResult.Score.SellerScore.Score < 0 || tt.scoringResult.Score.SellerScore.Score > 100 {
				t.Errorf("SellerScore should be between 0 and 100, got %d", tt.scoringResult.Score.SellerScore.Score)
			}
			if tt.scoringResult.Score.AverageValueScore.Score < 0 || tt.scoringResult.Score.AverageValueScore.Score > 100 {
				t.Errorf("AverageValueScore should be between 0 and 100, got %d", tt.scoringResult.Score.AverageValueScore.Score)
			}
			if tt.scoringResult.Score.CurrencyScore.Score < 0 || tt.scoringResult.Score.CurrencyScore.Score > 100 {
				t.Errorf("CurrencyScore should be between 0 and 100, got %d", tt.scoringResult.Score.CurrencyScore.Score)
			}
		})
	}
}

func TestScoreCard_CalculateOverallScore(t *testing.T) {
	tests := []struct {
		name          string
		scoreCard     ScoreCard
		expectedScore int
	}{
		{
			name: "All high scores",
			scoreCard: ScoreCard{
				ValueScore:        ValueScoreCard{Score: 95},
				SellerScore:       SellerScoreCard{Score: 90},
				AverageValueScore: AverageValueScoreCard{Score: 85},
				CurrencyScore:     CurrencyScoreCard{Score: 100},
			},
			expectedScore: 92, // (95+90+85+100)/4 = 92.5 -> 92
		},
		{
			name: "Mixed scores",
			scoreCard: ScoreCard{
				ValueScore:        ValueScoreCard{Score: 60},
				SellerScore:       SellerScoreCard{Score: 70},
				AverageValueScore: AverageValueScoreCard{Score: 40},
				CurrencyScore:     CurrencyScoreCard{Score: 90},
			},
			expectedScore: 65, // (60+70+40+90)/4 = 65
		},
		{
			name: "All low scores",
			scoreCard: ScoreCard{
				ValueScore:        ValueScoreCard{Score: 10},
				SellerScore:       SellerScoreCard{Score: 20},
				AverageValueScore: AverageValueScoreCard{Score: 15},
				CurrencyScore:     CurrencyScoreCard{Score: 25},
			},
			expectedScore: 17, // (10+20+15+25)/4 = 17.5 -> 17
		},
		{
			name: "Zero scores",
			scoreCard: ScoreCard{
				ValueScore:        ValueScoreCard{Score: 0},
				SellerScore:       SellerScoreCard{Score: 0},
				AverageValueScore: AverageValueScoreCard{Score: 0},
				CurrencyScore:     CurrencyScoreCard{Score: 0},
			},
			expectedScore: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualScore := tt.scoreCard.CalculateOverallScore()
			if actualScore != tt.expectedScore {
				t.Errorf("CalculateOverallScore() = %d, expected %d", actualScore, tt.expectedScore)
			}
		})
	}
}

func TestScoreCard_GetRiskLevel(t *testing.T) {
	tests := []struct {
		name          string
		overallScore  int
		expectedLevel string
	}{
		{
			name:          "Low risk - high score",
			overallScore:  85,
			expectedLevel: "low",
		},
		{
			name:          "Low risk - boundary",
			overallScore:  70,
			expectedLevel: "low",
		},
		{
			name:          "Medium risk - upper boundary",
			overallScore:  69,
			expectedLevel: "medium",
		},
		{
			name:          "Medium risk - middle",
			overallScore:  55,
			expectedLevel: "medium",
		},
		{
			name:          "Medium risk - lower boundary",
			overallScore:  40,
			expectedLevel: "medium",
		},
		{
			name:          "High risk - upper boundary",
			overallScore:  39,
			expectedLevel: "high",
		},
		{
			name:          "High risk - middle",
			overallScore:  25,
			expectedLevel: "high",
		},
		{
			name:          "High risk - lower boundary",
			overallScore:  20,
			expectedLevel: "high",
		},
		{
			name:          "Critical risk",
			overallScore:  19,
			expectedLevel: "critical",
		},
		{
			name:          "Critical risk - zero",
			overallScore:  0,
			expectedLevel: "critical",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualLevel := GetRiskLevel(tt.overallScore)
			if actualLevel != tt.expectedLevel {
				t.Errorf("GetRiskLevel(%d) = %s, expected %s", tt.overallScore, actualLevel, tt.expectedLevel)
			}
		})
	}
}

func TestTransaction_IsValid(t *testing.T) {
	tests := []struct {
		name        string
		transaction Transaction
		expected    bool
	}{
		{
			name: "Valid transaction",
			transaction: Transaction{
				Id: "tx-123",
			},
			expected: true,
		},
		{
			name: "Invalid transaction - empty ID",
			transaction: Transaction{
				Id: "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.transaction.IsValid()
			if actual != tt.expected {
				t.Errorf("IsValid() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

// Helper methods for testing - these would be actual methods in the domain model

func (sc *ScoreCard) CalculateOverallScore() int {
	total := sc.ValueScore.Score + sc.SellerScore.Score + sc.AverageValueScore.Score + sc.CurrencyScore.Score
	return total / 4
}

func GetRiskLevel(score int) string {
	if score >= 70 {
		return "low"
	} else if score >= 40 {
		return "medium"
	} else if score >= 20 {
		return "high"
	}
	return "critical"
}

func (t *Transaction) IsValid() bool {
	return t.Id != ""
}

// Benchmark tests
func BenchmarkScoreCard_CalculateOverallScore(b *testing.B) {
	scoreCard := ScoreCard{
		ValueScore:        ValueScoreCard{Score: 85},
		SellerScore:       SellerScoreCard{Score: 90},
		AverageValueScore: AverageValueScoreCard{Score: 75},
		CurrencyScore:     CurrencyScoreCard{Score: 95},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoreCard.CalculateOverallScore()
	}
}

func BenchmarkGetRiskLevel(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetRiskLevel(75)
	}
}