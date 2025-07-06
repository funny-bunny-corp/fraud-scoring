package domain

import (
	"reflect"
	"testing"
	"time"
)

// Helper function to create a valid transaction analysis for testing
func createValidTransactionForComponent() *TransactionAnalysis {
	return &TransactionAnalysis{
		Participants: Participants{
			Buyer: BuyerInfo{
				Document: "12345678901", // Valid 11-digit document
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
			At: time.Now().Add(-time.Hour), // 1 hour ago
		},
		Payment: Payment{
			Id:       "payment-789",
			Amount:   "100.00",
			Currency: "USD",
			Status:   "completed",
		},
	}
}

// Helper function to create a transaction component with default settings
func createDefaultTransactionComponent() *TransactionComponent {
	return NewTransactionComponent(
		24*time.Hour,                           // maxAge: 24 hours
		1.0,                                   // minAmount: $1.00
		10000.0,                               // maxAmount: $10,000.00
		[]string{"USD", "EUR", "BRL", "JPY"}, // supported currencies
	)
}

// Test Component function - Valid transaction
func TestTransactionComponent_Component_ValidTransaction(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if !result.IsValid {
		t.Errorf("Expected valid transaction, got invalid: %v", result.Errors)
	}

	if len(result.Errors) != 0 {
		t.Errorf("Expected no errors, got %v", result.Errors)
	}
}

// Test Component function - Nil transaction
func TestTransactionComponent_Component_NilTransaction(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()

	// Act
	result, err := tc.Component(nil)

	// Assert
	if err == nil {
		t.Error("Expected error for nil transaction")
	}

	if result != nil {
		t.Error("Expected nil result for nil transaction")
	}

	expectedError := "transaction cannot be nil"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

// Test Component function - Missing payment ID
func TestTransactionComponent_Component_MissingPaymentId(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Id = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "payment ID cannot be empty") {
		t.Errorf("Expected 'payment ID cannot be empty' error, got %v", result.Errors)
	}
}

// Test Component function - Missing order ID
func TestTransactionComponent_Component_MissingOrderId(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.Id = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "order ID cannot be empty") {
		t.Errorf("Expected 'order ID cannot be empty' error, got %v", result.Errors)
	}
}

// Test Component function - Invalid payment amount format
func TestTransactionComponent_Component_InvalidAmountFormat(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Amount = "invalid-amount"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "invalid payment amount format") {
		t.Errorf("Expected 'invalid payment amount format' error, got %v", result.Errors)
	}
}

// Test Component function - Amount below minimum threshold
func TestTransactionComponent_Component_AmountBelowMinimum(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Amount = "0.50" // Below minimum of $1.00

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "payment amount below minimum threshold") {
		t.Errorf("Expected 'payment amount below minimum threshold' error, got %v", result.Errors)
	}
}

// Test Component function - Amount above maximum threshold
func TestTransactionComponent_Component_AmountAboveMaximum(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Amount = "15000.00" // Above maximum of $10,000.00

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "payment amount exceeds maximum threshold") {
		t.Errorf("Expected 'payment amount exceeds maximum threshold' error, got %v", result.Errors)
	}
}

// Test Component function - Missing currency
func TestTransactionComponent_Component_MissingCurrency(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Currency = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "payment currency is required") {
		t.Errorf("Expected 'payment currency is required' error, got %v", result.Errors)
	}
}

// Test Component function - Unsupported currency
func TestTransactionComponent_Component_UnsupportedCurrency(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Currency = "XYZ"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "unsupported currency") {
		t.Errorf("Expected 'unsupported currency' error, got %v", result.Errors)
	}
}

// Test Component function - Invalid payment status
func TestTransactionComponent_Component_InvalidPaymentStatus(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Status = "invalid-status"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "invalid payment status") {
		t.Errorf("Expected 'invalid payment status' error, got %v", result.Errors)
	}
}

// Test Component function - Missing buyer document
func TestTransactionComponent_Component_MissingBuyerDocument(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Participants.Buyer.Document = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "buyer document is required") {
		t.Errorf("Expected 'buyer document is required' error, got %v", result.Errors)
	}
}

// Test Component function - Invalid buyer document format
func TestTransactionComponent_Component_InvalidBuyerDocument(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Participants.Buyer.Document = "invalid-doc"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "invalid buyer document format") {
		t.Errorf("Expected 'invalid buyer document format' error, got %v", result.Errors)
	}
}

// Test Component function - Missing buyer name
func TestTransactionComponent_Component_MissingBuyerName(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Participants.Buyer.Name = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "buyer name is required") {
		t.Errorf("Expected 'buyer name is required' error, got %v", result.Errors)
	}
}

// Test Component function - Buyer name too short
func TestTransactionComponent_Component_BuyerNameTooShort(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Participants.Buyer.Name = "J"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "buyer name too short") {
		t.Errorf("Expected 'buyer name too short' error, got %v", result.Errors)
	}
}

// Test Component function - Missing seller ID
func TestTransactionComponent_Component_MissingSellerId(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Participants.Seller.SellerId = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "seller ID is required") {
		t.Errorf("Expected 'seller ID is required' error, got %v", result.Errors)
	}
}

// Test Component function - Seller ID too short
func TestTransactionComponent_Component_SellerIdTooShort(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Participants.Seller.SellerId = "ab"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "seller ID too short") {
		t.Errorf("Expected 'seller ID too short' error, got %v", result.Errors)
	}
}

// Test Component function - Missing payment token
func TestTransactionComponent_Component_MissingPaymentToken(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.PaymentType.Token = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "payment token is required") {
		t.Errorf("Expected 'payment token is required' error, got %v", result.Errors)
	}
}

// Test Component function - Missing card info
func TestTransactionComponent_Component_MissingCardInfo(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.PaymentType.CardInfo = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "card info is required") {
		t.Errorf("Expected 'card info is required' error, got %v", result.Errors)
	}
}

// Test Component function - Invalid card info format
func TestTransactionComponent_Component_InvalidCardInfo(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.PaymentType.CardInfo = "invalid-card"

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "invalid card info format") {
		t.Errorf("Expected 'invalid card info format' error, got %v", result.Errors)
	}
}

// Test Component function - Transaction too old
func TestTransactionComponent_Component_TransactionTooOld(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.At = time.Now().Add(-48 * time.Hour) // 48 hours ago (older than 24h limit)

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "transaction too old") {
		t.Errorf("Expected 'transaction too old' error, got %v", result.Errors)
	}
}

// Test Component function - Transaction in future
func TestTransactionComponent_Component_TransactionInFuture(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.At = time.Now().Add(2 * time.Hour) // 2 hours in future

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "transaction timestamp is in the future") {
		t.Errorf("Expected 'transaction timestamp is in the future' error, got %v", result.Errors)
	}
}

// Test Component function - Zero timestamp
func TestTransactionComponent_Component_ZeroTimestamp(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Order.At = time.Time{} // Zero timestamp

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	if !containsError(result.Errors, "checkout timestamp is required") {
		t.Errorf("Expected 'checkout timestamp is required' error, got %v", result.Errors)
	}
}

// Test Component function - Multiple validation errors
func TestTransactionComponent_Component_MultipleErrors(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	
	// Set multiple invalid fields
	transaction.Payment.Id = ""
	transaction.Order.Id = ""
	transaction.Payment.Amount = ""
	transaction.Payment.Currency = ""
	transaction.Participants.Buyer.Document = ""
	transaction.Participants.Seller.SellerId = ""

	// Act
	result, err := tc.Component(transaction)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.IsValid {
		t.Error("Expected invalid transaction")
	}

	expectedErrors := []string{
		"payment ID cannot be empty",
		"payment ID is required",
		"buyer document is required",
		"checkout ID is required",
	}

	for _, expectedError := range expectedErrors {
		if !containsError(result.Errors, expectedError) {
			t.Errorf("Expected error '%s' not found in %v", expectedError, result.Errors)
		}
	}
}

// Test Component function - Valid card info variations
func TestTransactionComponent_Component_ValidCardInfoVariations(t *testing.T) {
	testCases := []struct {
		name     string
		cardInfo string
	}{
		{"Masked card format", "****1234"},
		{"Full card number", "1234567890123456"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			component := createDefaultTransactionComponent()
			transaction := createValidTransactionForComponent()
			transaction.Order.PaymentType.CardInfo = tc.cardInfo

			// Act
			result, err := component.Component(transaction)

			// Assert
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if !result.IsValid {
				t.Errorf("Expected valid transaction for card info '%s', got errors: %v", tc.cardInfo, result.Errors)
			}
		})
	}
}

// Test Component function - Different currencies
func TestTransactionComponent_Component_DifferentCurrencies(t *testing.T) {
	currencies := []string{"USD", "EUR", "BRL", "JPY"}

	for _, currency := range currencies {
		t.Run("Currency_"+currency, func(t *testing.T) {
			// Arrange
			tc := createDefaultTransactionComponent()
			transaction := createValidTransactionForComponent()
			transaction.Payment.Currency = currency

			// Act
			result, err := tc.Component(transaction)

			// Assert
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if !result.IsValid {
				t.Errorf("Expected valid transaction for currency %s, got errors: %v", currency, result.Errors)
			}
		})
	}
}

// Test Component function - Different payment statuses
func TestTransactionComponent_Component_DifferentPaymentStatuses(t *testing.T) {
	statuses := []string{"pending", "completed", "failed", "cancelled"}

	for _, status := range statuses {
		t.Run("Status_"+status, func(t *testing.T) {
			// Arrange
			tc := createDefaultTransactionComponent()
			transaction := createValidTransactionForComponent()
			transaction.Payment.Status = status

			// Act
			result, err := tc.Component(transaction)

			// Assert
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if !result.IsValid {
				t.Errorf("Expected valid transaction for status %s, got errors: %v", status, result.Errors)
			}
		})
	}
}

// Test Component function - Edge case amounts
func TestTransactionComponent_Component_EdgeCaseAmounts(t *testing.T) {
	testCases := []struct {
		name     string
		amount   string
		expected bool
	}{
		{"Minimum amount", "1.00", true},
		{"Maximum amount", "10000.00", true},
		{"Just below minimum", "0.99", false},
		{"Just above maximum", "10000.01", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			component := createDefaultTransactionComponent()
			transaction := createValidTransactionForComponent()
			transaction.Payment.Amount = tc.amount

			// Act
			result, err := component.Component(transaction)

			// Assert
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if result.IsValid != tc.expected {
				t.Errorf("Expected IsValid=%v for amount %s, got %v", tc.expected, tc.amount, result.IsValid)
			}
		})
	}
}

// Test NewTransactionComponent
func TestNewTransactionComponent(t *testing.T) {
	// Arrange
	maxAge := 24 * time.Hour
	minAmount := 1.0
	maxAmount := 10000.0
	currencies := []string{"USD", "EUR"}

	// Act
	tc := NewTransactionComponent(maxAge, minAmount, maxAmount, currencies)

	// Assert
	if tc == nil {
		t.Fatal("Expected TransactionComponent instance, got nil")
	}

	if tc.maxTransactionAge != maxAge {
		t.Errorf("Expected maxTransactionAge %v, got %v", maxAge, tc.maxTransactionAge)
	}

	if tc.minAmount != minAmount {
		t.Errorf("Expected minAmount %v, got %v", minAmount, tc.minAmount)
	}

	if tc.maxAmount != maxAmount {
		t.Errorf("Expected maxAmount %v, got %v", maxAmount, tc.maxAmount)
	}

	if !reflect.DeepEqual(tc.supportedCurrencies, currencies) {
		t.Errorf("Expected supportedCurrencies %v, got %v", currencies, tc.supportedCurrencies)
	}
}

// Test GetSupportedCurrencies
func TestTransactionComponent_GetSupportedCurrencies(t *testing.T) {
	// Arrange
	currencies := []string{"USD", "EUR", "BRL", "JPY"}
	tc := NewTransactionComponent(time.Hour, 1.0, 1000.0, currencies)

	// Act
	result := tc.GetSupportedCurrencies()

	// Assert
	if !reflect.DeepEqual(result, currencies) {
		t.Errorf("Expected currencies %v, got %v", currencies, result)
	}
}

// Test GetLimits
func TestTransactionComponent_GetLimits(t *testing.T) {
	// Arrange
	minAmount := 5.0
	maxAmount := 2000.0
	tc := NewTransactionComponent(time.Hour, minAmount, maxAmount, []string{"USD"})

	// Act
	min, max := tc.GetLimits()

	// Assert
	if min != minAmount {
		t.Errorf("Expected min amount %v, got %v", minAmount, min)
	}

	if max != maxAmount {
		t.Errorf("Expected max amount %v, got %v", maxAmount, max)
	}
}

// Test GetMaxAge
func TestTransactionComponent_GetMaxAge(t *testing.T) {
	// Arrange
	maxAge := 48 * time.Hour
	tc := NewTransactionComponent(maxAge, 1.0, 1000.0, []string{"USD"})

	// Act
	result := tc.GetMaxAge()

	// Assert
	if result != maxAge {
		t.Errorf("Expected max age %v, got %v", maxAge, result)
	}
}

// Test isValidDocument edge cases
func TestTransactionComponent_isValidDocument_EdgeCases(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	
	testCases := []struct {
		name     string
		document string
		expected bool
	}{
		{"Valid document", "12345678901", true},
		{"Too short", "1234567890", false},
		{"Too long", "123456789012", false},
		{"Contains letters", "1234567890a", false},
		{"Contains special chars", "1234567890!", false},
		{"Empty string", "", false},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Act
			result := tc.isValidDocument(test.document)

			// Assert
			if result != test.expected {
				t.Errorf("Expected %v for document '%s', got %v", test.expected, test.document, result)
			}
		})
	}
}

// Test isValidCardInfo edge cases
func TestTransactionComponent_isValidCardInfo_EdgeCases(t *testing.T) {
	// Arrange
	tc := createDefaultTransactionComponent()
	
	testCases := []struct {
		name     string
		cardInfo string
		expected bool
	}{
		{"Valid masked format", "****1234", true},
		{"Valid full format", "1234567890123456", true},
		{"Invalid masked - wrong stars", "***1234", false},
		{"Invalid masked - wrong digits", "****123", false},
		{"Invalid full - too short", "123456789012345", false},
		{"Invalid full - too long", "12345678901234567", false},
		{"Invalid full - contains letters", "123456789012345a", false},
		{"Empty string", "", false},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Act
			result := tc.isValidCardInfo(test.cardInfo)

			// Assert
			if result != test.expected {
				t.Errorf("Expected %v for card info '%s', got %v", test.expected, test.cardInfo, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkTransactionComponent_Component_Valid(b *testing.B) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()

	// Act
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tc.Component(transaction)
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkTransactionComponent_Component_Invalid(b *testing.B) {
	// Arrange
	tc := createDefaultTransactionComponent()
	transaction := createValidTransactionForComponent()
	transaction.Payment.Id = "" // Make it invalid

	// Act
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tc.Component(transaction)
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkNewTransactionComponent(b *testing.B) {
	// Arrange
	maxAge := 24 * time.Hour
	minAmount := 1.0
	maxAmount := 10000.0
	currencies := []string{"USD", "EUR", "BRL", "JPY"}

	// Act
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewTransactionComponent(maxAge, minAmount, maxAmount, currencies)
	}
}

// Helper function to check if errors slice contains a specific error message
func containsError(errors []string, target string) bool {
	for _, err := range errors {
		if err == target {
			return true
		}
	}
	return false
}