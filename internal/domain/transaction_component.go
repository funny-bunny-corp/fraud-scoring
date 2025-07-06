package domain

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

// TransactionComponent provides transaction validation and processing functionality
type TransactionComponent struct {
	maxTransactionAge time.Duration
	minAmount         float64
	maxAmount         float64
	supportedCurrencies []string
}

// NewTransactionComponent creates a new instance of TransactionComponent
func NewTransactionComponent(maxAge time.Duration, minAmount, maxAmount float64, currencies []string) *TransactionComponent {
	return &TransactionComponent{
		maxTransactionAge:   maxAge,
		minAmount:          minAmount,
		maxAmount:          maxAmount,
		supportedCurrencies: currencies,
	}
}

// ComponentValidationResult represents the result of component validation
type ComponentValidationResult struct {
	IsValid bool
	Errors  []string
}

// Component validates a transaction analysis and returns validation results
func (tc *TransactionComponent) Component(transaction *TransactionAnalysis) (*ComponentValidationResult, error) {
	if transaction == nil {
		return nil, errors.New("transaction cannot be nil")
	}

	result := &ComponentValidationResult{
		IsValid: true,
		Errors:  []string{},
	}

	// Validate transaction structure
	if err := tc.validateTransactionStructure(transaction); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	// Validate payment information
	if err := tc.validatePayment(&transaction.Payment); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	// Validate participants
	if err := tc.validateParticipants(&transaction.Participants); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	// Validate checkout information
	if err := tc.validateCheckout(&transaction.Order); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	// Validate transaction age
	if err := tc.validateTransactionAge(&transaction.Order); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	return result, nil
}

// validateTransactionStructure validates the basic structure of the transaction
func (tc *TransactionComponent) validateTransactionStructure(transaction *TransactionAnalysis) error {
	if transaction.Payment.Id == "" {
		return errors.New("payment ID cannot be empty")
	}
	if transaction.Order.Id == "" {
		return errors.New("order ID cannot be empty")
	}
	return nil
}

// validatePayment validates payment information
func (tc *TransactionComponent) validatePayment(payment *Payment) error {
	if payment.Id == "" {
		return errors.New("payment ID is required")
	}

	if payment.Amount == "" {
		return errors.New("payment amount is required")
	}

	amount, err := strconv.ParseFloat(payment.Amount, 64)
	if err != nil {
		return errors.New("invalid payment amount format")
	}

	if amount < tc.minAmount {
		return errors.New("payment amount below minimum threshold")
	}

	if amount > tc.maxAmount {
		return errors.New("payment amount exceeds maximum threshold")
	}

	if payment.Currency == "" {
		return errors.New("payment currency is required")
	}

	if !tc.isSupportedCurrency(payment.Currency) {
		return errors.New("unsupported currency")
	}

	validStatuses := []string{"pending", "completed", "failed", "cancelled"}
	if !tc.isValidStatus(payment.Status, validStatuses) {
		return errors.New("invalid payment status")
	}

	return nil
}

// validateParticipants validates buyer and seller information
func (tc *TransactionComponent) validateParticipants(participants *Participants) error {
	// Validate buyer
	if participants.Buyer.Document == "" {
		return errors.New("buyer document is required")
	}

	if !tc.isValidDocument(participants.Buyer.Document) {
		return errors.New("invalid buyer document format")
	}

	if participants.Buyer.Name == "" {
		return errors.New("buyer name is required")
	}

	if len(participants.Buyer.Name) < 2 {
		return errors.New("buyer name too short")
	}

	// Validate seller
	if participants.Seller.SellerId == "" {
		return errors.New("seller ID is required")
	}

	if len(participants.Seller.SellerId) < 3 {
		return errors.New("seller ID too short")
	}

	return nil
}

// validateCheckout validates checkout information
func (tc *TransactionComponent) validateCheckout(checkout *Checkout) error {
	if checkout.Id == "" {
		return errors.New("checkout ID is required")
	}

	if checkout.PaymentType.Token == "" {
		return errors.New("payment token is required")
	}

	if checkout.PaymentType.CardInfo == "" {
		return errors.New("card info is required")
	}

	if !tc.isValidCardInfo(checkout.PaymentType.CardInfo) {
		return errors.New("invalid card info format")
	}

	return nil
}

// validateTransactionAge validates if transaction is within acceptable age limit
func (tc *TransactionComponent) validateTransactionAge(checkout *Checkout) error {
	if checkout.At.IsZero() {
		return errors.New("checkout timestamp is required")
	}

	age := time.Since(checkout.At)
	if age > tc.maxTransactionAge {
		return errors.New("transaction too old")
	}

	if checkout.At.After(time.Now().Add(time.Hour)) {
		return errors.New("transaction timestamp is in the future")
	}

	return nil
}

// Helper functions

func (tc *TransactionComponent) isSupportedCurrency(currency string) bool {
	for _, supported := range tc.supportedCurrencies {
		if supported == currency {
			return true
		}
	}
	return false
}

func (tc *TransactionComponent) isValidStatus(status string, validStatuses []string) bool {
	for _, validStatus := range validStatuses {
		if validStatus == status {
			return true
		}
	}
	return false
}

func (tc *TransactionComponent) isValidDocument(document string) bool {
	// Simple validation for document format (digits only, 11 characters for CPF)
	if len(document) != 11 {
		return false
	}
	
	matched, _ := regexp.MatchString(`^\d{11}$`, document)
	return matched
}

func (tc *TransactionComponent) isValidCardInfo(cardInfo string) bool {
	// Check for masked card format (****1234)
	masked, _ := regexp.MatchString(`^\*{4}\d{4}$`, cardInfo)
	if masked {
		return true
	}
	
	// Check for full card number format (16 digits)
	full, _ := regexp.MatchString(`^\d{16}$`, cardInfo)
	return full
}

// GetSupportedCurrencies returns the list of supported currencies
func (tc *TransactionComponent) GetSupportedCurrencies() []string {
	return tc.supportedCurrencies
}

// GetLimits returns the transaction limits
func (tc *TransactionComponent) GetLimits() (float64, float64) {
	return tc.minAmount, tc.maxAmount
}

// GetMaxAge returns the maximum transaction age
func (tc *TransactionComponent) GetMaxAge() time.Duration {
	return tc.maxTransactionAge
}