package application

import (
	"fraud-scoring/internal/domain"
	"fraud-scoring/internal/domain/application/errors"
	"fraud-scoring/internal/domain/repositories"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// Mock implementations for testing
type mockUserTransactionsRepository struct {
	lastOrderFunc           func(string) (*domain.TransactionAnalysis, error)
	averageTransactionsFunc func(string, time.Time) (*domain.UserMonthAverage, error)
}

func (m *mockUserTransactionsRepository) LastOrder(document string) (*domain.TransactionAnalysis, error) {
	if m.lastOrderFunc != nil {
		return m.lastOrderFunc(document)
	}
	return nil, nil
}

func (m *mockUserTransactionsRepository) AverageTransactions(document string, date time.Time) (*domain.UserMonthAverage, error) {
	if m.averageTransactionsFunc != nil {
		return m.averageTransactionsFunc(document, date)
	}
	return nil, nil
}

type mockTransactionScoreCard struct {
	storeFunc func(*domain.ScoringResult) error
}

func (m *mockTransactionScoreCard) Store(scoreCard *domain.ScoringResult) error {
	if m.storeFunc != nil {
		return m.storeFunc(scoreCard)
	}
	return nil
}

// Helper function to create a valid transaction analysis
func createValidTransactionAnalysis() *domain.TransactionAnalysis {
	return &domain.TransactionAnalysis{
		Participants: domain.Participants{
			Buyer: domain.BuyerInfo{
				Document: "12345678901",
				Name:     "John Doe",
			},
			Seller: domain.SellerInfo{
				SellerId: "seller-123",
			},
		},
		Order: domain.Checkout{
			Id: "order-456",
			PaymentType: domain.CardInfo{
				CardInfo: "****1234",
				Token:    "tok_123",
			},
			At: time.Now(),
		},
		Payment: domain.Payment{
			Id:       "payment-789",
			Amount:   "100.00",
			Currency: "USD",
			Status:   "completed",
		},
	}
}

func TestPaymentRiskScoring_Assessment_Success(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	mockUTR := &mockUserTransactionsRepository{
		lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
			return createValidTransactionAnalysis(), nil
		},
		averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
			return &domain.UserMonthAverage{
				Document: document,
				Month:    "2024-01",
				Total:    "1000.00",
				Count:    10,
			}, nil
		},
	}

	mockTSC := &mockTransactionScoreCard{
		storeFunc: func(scoreCard *domain.ScoringResult) error {
			return nil
		},
	}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
	transaction := createValidTransactionAnalysis()

	err := prs.Assessment(transaction)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPaymentRiskScoring_Assessment_LastOrderError(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	mockUTR := &mockUserTransactionsRepository{
		lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
			return nil, errors.New("database error")
		},
	}

	mockTSC := &mockTransactionScoreCard{}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
	transaction := createValidTransactionAnalysis()

	err := prs.Assessment(transaction)

	if err == nil {
		t.Error("Expected error, got none")
	}

	if _, ok := err.(errors.LastOrderNotFound); !ok {
		t.Errorf("Expected LastOrderNotFound error, got %T", err)
	}
}

func TestPaymentRiskScoring_Assessment_AverageTransactionsError(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	mockUTR := &mockUserTransactionsRepository{
		lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
			return createValidTransactionAnalysis(), nil
		},
		averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
			return nil, errors.New("database error")
		},
	}

	mockTSC := &mockTransactionScoreCard{}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
	transaction := createValidTransactionAnalysis()

	err := prs.Assessment(transaction)

	if err == nil {
		t.Error("Expected error, got none")
	}

	if _, ok := err.(errors.AverageTransactionsNotFound); !ok {
		t.Errorf("Expected AverageTransactionsNotFound error, got %T", err)
	}
}

func TestPaymentRiskScoring_Assessment_StoreError(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	mockUTR := &mockUserTransactionsRepository{
		lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
			return createValidTransactionAnalysis(), nil
		},
		averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
			return &domain.UserMonthAverage{
				Document: document,
				Month:    "2024-01",
				Total:    "1000.00",
				Count:    10,
			}, nil
		},
	}

	mockTSC := &mockTransactionScoreCard{
		storeFunc: func(scoreCard *domain.ScoringResult) error {
			return errors.New("storage error")
		},
	}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
	transaction := createValidTransactionAnalysis()

	err := prs.Assessment(transaction)

	if err == nil {
		t.Error("Expected error, got none")
	}

	if err.Error() != "storage error" {
		t.Errorf("Expected 'storage error', got %v", err)
	}
}

func TestPaymentRiskScoring_Assessment_NilTransaction(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	mockUTR := &mockUserTransactionsRepository{}
	mockTSC := &mockTransactionScoreCard{}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)

	// This should panic or handle nil gracefully
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for nil transaction")
		}
	}()

	err := prs.Assessment(nil)
	if err == nil {
		t.Error("Expected error for nil transaction")
	}
}

func TestPaymentRiskScoring_Assessment_ScoreCardGeneration(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	var storedScoreCard *domain.ScoringResult

	mockUTR := &mockUserTransactionsRepository{
		lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
			return createValidTransactionAnalysis(), nil
		},
		averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
			return &domain.UserMonthAverage{
				Document: document,
				Month:    "2024-01",
				Total:    "1000.00",
				Count:    10,
			}, nil
		},
	}

	mockTSC := &mockTransactionScoreCard{
		storeFunc: func(scoreCard *domain.ScoringResult) error {
			storedScoreCard = scoreCard
			return nil
		},
	}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
	transaction := createValidTransactionAnalysis()

	err := prs.Assessment(transaction)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if storedScoreCard == nil {
		t.Error("Expected scoreCard to be stored")
		return
	}

	// Verify scoreCard structure
	if storedScoreCard.Transaction.Payment.Id != transaction.Payment.Id {
		t.Errorf("Expected transaction ID %s, got %s", 
			transaction.Payment.Id, storedScoreCard.Transaction.Payment.Id)
	}

	if storedScoreCard.Transaction.Participants.Buyer.Document != transaction.Participants.Buyer.Document {
		t.Errorf("Expected buyer document %s, got %s", 
			transaction.Participants.Buyer.Document, storedScoreCard.Transaction.Participants.Buyer.Document)
	}

	// Verify that scores are within valid range (0-100)
	scores := []int{
		storedScoreCard.Score.ValueScore.Score,
		storedScoreCard.Score.SellerScore.Score,
		storedScoreCard.Score.AverageValueScore.Score,
		storedScoreCard.Score.CurrencyScore.Score,
	}

	for i, score := range scores {
		if score < 0 || score > 100 {
			t.Errorf("Score %d is out of range [0-100]: %d", i, score)
		}
	}
}

func TestNewPaymentRiskScoring(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	mockUTR := &mockUserTransactionsRepository{}
	mockTSC := &mockTransactionScoreCard{}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)

	if prs == nil {
		t.Error("Expected PaymentRiskScoring instance, got nil")
	}

	if prs.utr != mockUTR {
		t.Error("Expected user transactions repository to be set")
	}

	if prs.tsc != mockTSC {
		t.Error("Expected transaction score card repository to be set")
	}

	if prs.log != logger {
		t.Error("Expected logger to be set")
	}
}

func TestPaymentRiskScoring_Assessment_WithDifferentCurrencies(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	currencies := []string{"USD", "EUR", "BRL", "JPY"}

	for _, currency := range currencies {
		t.Run("Currency_"+currency, func(t *testing.T) {
			mockUTR := &mockUserTransactionsRepository{
				lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
					return createValidTransactionAnalysis(), nil
				},
				averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
					return &domain.UserMonthAverage{
						Document: document,
						Month:    "2024-01",
						Total:    "1000.00",
						Count:    10,
					}, nil
				},
			}

			mockTSC := &mockTransactionScoreCard{
				storeFunc: func(scoreCard *domain.ScoringResult) error {
					return nil
				},
			}

			prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
			transaction := createValidTransactionAnalysis()
			transaction.Payment.Currency = currency

			err := prs.Assessment(transaction)

			if err != nil {
				t.Errorf("Expected no error for currency %s, got %v", currency, err)
			}
		})
	}
}

func TestPaymentRiskScoring_Assessment_WithDifferentAmounts(t *testing.T) {
	logger := zaptest.NewLogger(t)
	defer logger.Sync()

	amounts := []string{"10.00", "100.00", "1000.00", "10000.00"}

	for _, amount := range amounts {
		t.Run("Amount_"+amount, func(t *testing.T) {
			mockUTR := &mockUserTransactionsRepository{
				lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
					return createValidTransactionAnalysis(), nil
				},
				averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
					return &domain.UserMonthAverage{
						Document: document,
						Month:    "2024-01",
						Total:    "1000.00",
						Count:    10,
					}, nil
				},
			}

			mockTSC := &mockTransactionScoreCard{
				storeFunc: func(scoreCard *domain.ScoringResult) error {
					return nil
				},
			}

			prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
			transaction := createValidTransactionAnalysis()
			transaction.Payment.Amount = amount

			err := prs.Assessment(transaction)

			if err != nil {
				t.Errorf("Expected no error for amount %s, got %v", amount, err)
			}
		})
	}
}

// Benchmark tests
func BenchmarkPaymentRiskScoring_Assessment(b *testing.B) {
	logger := zap.NewNop()

	mockUTR := &mockUserTransactionsRepository{
		lastOrderFunc: func(document string) (*domain.TransactionAnalysis, error) {
			return createValidTransactionAnalysis(), nil
		},
		averageTransactionsFunc: func(document string, date time.Time) (*domain.UserMonthAverage, error) {
			return &domain.UserMonthAverage{
				Document: document,
				Month:    "2024-01",
				Total:    "1000.00",
				Count:    10,
			}, nil
		},
	}

	mockTSC := &mockTransactionScoreCard{
		storeFunc: func(scoreCard *domain.ScoringResult) error {
			return nil
		},
	}

	prs := NewPaymentRiskScoring(mockUTR, mockTSC, logger)
	transaction := createValidTransactionAnalysis()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := prs.Assessment(transaction)
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

// Additional required types and functions for testing
type errors struct{}

func (e errors) New(message string) error {
	return &customError{message: message}
}

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

// Additional domain types that might be missing
type UserMonthAverage struct {
	Document string
	Month    string
	Total    string
	Count    int
}

// Add this to domain package if not exists
func (d *domain) UserMonthAverage() *UserMonthAverage {
	return &UserMonthAverage{}
}