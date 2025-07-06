package domain

import (
	"testing"
	"time"
)

func TestTransactionAnalysis_IsValid(t *testing.T) {
	validTime := time.Now()
	
	tests := []struct {
		name        string
		transaction TransactionAnalysis
		expected    bool
	}{
		{
			name: "Valid transaction analysis",
			transaction: TransactionAnalysis{
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
					At: validTime,
				},
				Payment: Payment{
					Id:       "payment-789",
					Amount:   "100.00",
					Currency: "USD",
					Status:   "completed",
				},
			},
			expected: true,
		},
		{
			name: "Invalid transaction analysis - missing buyer document",
			transaction: TransactionAnalysis{
				Participants: Participants{
					Buyer: BuyerInfo{
						Document: "",
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
					At: validTime,
				},
				Payment: Payment{
					Id:       "payment-789",
					Amount:   "100.00",
					Currency: "USD",
					Status:   "completed",
				},
			},
			expected: false,
		},
		{
			name: "Invalid transaction analysis - missing seller ID",
			transaction: TransactionAnalysis{
				Participants: Participants{
					Buyer: BuyerInfo{
						Document: "12345678901",
						Name:     "John Doe",
					},
					Seller: SellerInfo{
						SellerId: "",
					},
				},
				Order: Checkout{
					Id: "order-456",
					PaymentType: CardInfo{
						CardInfo: "****1234",
						Token:    "tok_123",
					},
					At: validTime,
				},
				Payment: Payment{
					Id:       "payment-789",
					Amount:   "100.00",
					Currency: "USD",
					Status:   "completed",
				},
			},
			expected: false,
		},
		{
			name: "Invalid transaction analysis - missing payment ID",
			transaction: TransactionAnalysis{
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
					At: validTime,
				},
				Payment: Payment{
					Id:       "",
					Amount:   "100.00",
					Currency: "USD",
					Status:   "completed",
				},
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

func TestPayment_IsCompleted(t *testing.T) {
	tests := []struct {
		name     string
		payment  Payment
		expected bool
	}{
		{
			name: "Completed payment",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "100.00",
				Currency: "USD",
				Status:   "completed",
			},
			expected: true,
		},
		{
			name: "Pending payment",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "100.00",
				Currency: "USD",
				Status:   "pending",
			},
			expected: false,
		},
		{
			name: "Failed payment",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "100.00",
				Currency: "USD",
				Status:   "failed",
			},
			expected: false,
		},
		{
			name: "Cancelled payment",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "100.00",
				Currency: "USD",
				Status:   "cancelled",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payment.IsCompleted()
			if actual != tt.expected {
				t.Errorf("IsCompleted() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestPayment_GetAmountFloat(t *testing.T) {
	tests := []struct {
		name          string
		payment       Payment
		expectedValue float64
		expectedError bool
	}{
		{
			name: "Valid amount",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "100.50",
				Currency: "USD",
				Status:   "completed",
			},
			expectedValue: 100.50,
			expectedError: false,
		},
		{
			name: "Zero amount",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "0.00",
				Currency: "USD",
				Status:   "completed",
			},
			expectedValue: 0.00,
			expectedError: false,
		},
		{
			name: "Invalid amount format",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "invalid",
				Currency: "USD",
				Status:   "completed",
			},
			expectedValue: 0.00,
			expectedError: true,
		},
		{
			name: "Empty amount",
			payment: Payment{
				Id:       "payment-123",
				Amount:   "",
				Currency: "USD",
				Status:   "completed",
			},
			expectedValue: 0.00,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue, actualError := tt.payment.GetAmountFloat()
			
			if tt.expectedError && actualError == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectedError && actualError != nil {
				t.Errorf("Expected no error but got: %v", actualError)
			}
			if actualValue != tt.expectedValue {
				t.Errorf("GetAmountFloat() = %f, expected %f", actualValue, tt.expectedValue)
			}
		})
	}
}

func TestBuyerInfo_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		buyer    BuyerInfo
		expected bool
	}{
		{
			name: "Valid buyer",
			buyer: BuyerInfo{
				Document: "12345678901",
				Name:     "John Doe",
			},
			expected: true,
		},
		{
			name: "Invalid buyer - empty document",
			buyer: BuyerInfo{
				Document: "",
				Name:     "John Doe",
			},
			expected: false,
		},
		{
			name: "Invalid buyer - empty name",
			buyer: BuyerInfo{
				Document: "12345678901",
				Name:     "",
			},
			expected: false,
		},
		{
			name: "Invalid buyer - both empty",
			buyer: BuyerInfo{
				Document: "",
				Name:     "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.buyer.IsValid()
			if actual != tt.expected {
				t.Errorf("IsValid() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestSellerInfo_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		seller   SellerInfo
		expected bool
	}{
		{
			name: "Valid seller",
			seller: SellerInfo{
				SellerId: "seller-123",
			},
			expected: true,
		},
		{
			name: "Invalid seller - empty ID",
			seller: SellerInfo{
				SellerId: "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.seller.IsValid()
			if actual != tt.expected {
				t.Errorf("IsValid() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestCheckout_IsRecent(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name     string
		checkout Checkout
		duration time.Duration
		expected bool
	}{
		{
			name: "Recent checkout - 5 minutes ago",
			checkout: Checkout{
				Id: "checkout-123",
				PaymentType: CardInfo{
					CardInfo: "****1234",
					Token:    "tok_123",
				},
				At: now.Add(-5 * time.Minute),
			},
			duration: 10 * time.Minute,
			expected: true,
		},
		{
			name: "Not recent checkout - 15 minutes ago",
			checkout: Checkout{
				Id: "checkout-123",
				PaymentType: CardInfo{
					CardInfo: "****1234",
					Token:    "tok_123",
				},
				At: now.Add(-15 * time.Minute),
			},
			duration: 10 * time.Minute,
			expected: false,
		},
		{
			name: "Future checkout",
			checkout: Checkout{
				Id: "checkout-123",
				PaymentType: CardInfo{
					CardInfo: "****1234",
					Token:    "tok_123",
				},
				At: now.Add(5 * time.Minute),
			},
			duration: 10 * time.Minute,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.checkout.IsRecent(tt.duration)
			if actual != tt.expected {
				t.Errorf("IsRecent() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestCardInfo_IsMasked(t *testing.T) {
	tests := []struct {
		name     string
		cardInfo CardInfo
		expected bool
	}{
		{
			name: "Masked card info",
			cardInfo: CardInfo{
				CardInfo: "****1234",
				Token:    "tok_123",
			},
			expected: true,
		},
		{
			name: "Not masked card info",
			cardInfo: CardInfo{
				CardInfo: "1234567890123456",
				Token:    "tok_123",
			},
			expected: false,
		},
		{
			name: "Partially masked card info",
			cardInfo: CardInfo{
				CardInfo: "1234****1234",
				Token:    "tok_123",
			},
			expected: true,
		},
		{
			name: "Empty card info",
			cardInfo: CardInfo{
				CardInfo: "",
				Token:    "tok_123",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.cardInfo.IsMasked()
			if actual != tt.expected {
				t.Errorf("IsMasked() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

// Helper methods for testing - these would be actual methods in the domain model

func (ta *TransactionAnalysis) IsValid() bool {
	return ta.Participants.Buyer.IsValid() && 
		   ta.Participants.Seller.IsValid() && 
		   ta.Payment.Id != ""
}

func (p *Payment) IsCompleted() bool {
	return p.Status == "completed"
}

func (p *Payment) GetAmountFloat() (float64, error) {
	if p.Amount == "" {
		return 0.0, fmt.Errorf("amount is empty")
	}
	
	value, err := strconv.ParseFloat(p.Amount, 64)
	if err != nil {
		return 0.0, fmt.Errorf("invalid amount format: %v", err)
	}
	
	return value, nil
}

func (b *BuyerInfo) IsValid() bool {
	return b.Document != "" && b.Name != ""
}

func (s *SellerInfo) IsValid() bool {
	return s.SellerId != ""
}

func (c *Checkout) IsRecent(duration time.Duration) bool {
	return time.Since(c.At) <= duration
}

func (ci *CardInfo) IsMasked() bool {
	return strings.Contains(ci.CardInfo, "*")
}

// Additional required imports
import (
	"fmt"
	"strconv"
	"strings"
)

// Benchmark tests
func BenchmarkTransactionAnalysis_IsValid(b *testing.B) {
	validTime := time.Now()
	transaction := TransactionAnalysis{
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
			At: validTime,
		},
		Payment: Payment{
			Id:       "payment-789",
			Amount:   "100.00",
			Currency: "USD",
			Status:   "completed",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transaction.IsValid()
	}
}

func BenchmarkPayment_GetAmountFloat(b *testing.B) {
	payment := Payment{
		Id:       "payment-123",
		Amount:   "100.50",
		Currency: "USD",
		Status:   "completed",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payment.GetAmountFloat()
	}
}