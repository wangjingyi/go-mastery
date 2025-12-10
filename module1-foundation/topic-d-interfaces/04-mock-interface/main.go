// Assignment 4: The Mock Interface
//
// Goal: Define PaymentProcessor interface.
//       Create StripeProcessor (real) and MockProcessor (fake).
//       Swap them in main.
//
// Instructions:
// 1. Define PaymentProcessor interface with Charge method
// 2. Create StripeProcessor that would call real Stripe API
// 3. Create MockProcessor for testing
// 4. Demonstrate swapping implementations
//
// Key insight: This is how you write testable code in Go!

package main

import (
	"errors"
	"fmt"
)

// PaymentProcessor interface - defines the contract
type PaymentProcessor interface {
	Charge(amount float64, currency string) (transactionID string, err error)
	Refund(transactionID string) error
}

// StripeProcessor - the REAL implementation
type StripeProcessor struct {
	APIKey string
}

func (s *StripeProcessor) Charge(amount float64, currency string) (string, error) {
	// In real code, this would call Stripe's API
	fmt.Printf("[STRIPE] Charging $%.2f %s via Stripe API...\n", amount, currency)
	return "txn_stripe_12345", nil
}

func (s *StripeProcessor) Refund(transactionID string) error {
	fmt.Printf("[STRIPE] Refunding transaction %s...\n", transactionID)
	return nil
}

// MockProcessor - for TESTING
type MockProcessor struct {
	ShouldFail  bool
	ChargeCount int
}

func (m *MockProcessor) Charge(amount float64, currency string) (string, error) {
	m.ChargeCount++
	if m.ShouldFail {
		return "", errors.New("mock payment failed")
	}
	fmt.Printf("[MOCK] Pretending to charge $%.2f %s\n", amount, currency)
	return "txn_mock_99999", nil
}

func (m *MockProcessor) Refund(transactionID string) error {
	if m.ShouldFail {
		return errors.New("mock refund failed")
	}
	fmt.Printf("[MOCK] Pretending to refund %s\n", transactionID)
	return nil
}

// ProcessOrder uses the PaymentProcessor interface
// It doesn't know or care which implementation it's using!
func ProcessOrder(processor PaymentProcessor, amount float64) error {
	txnID, err := processor.Charge(amount, "USD")
	if err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}
	fmt.Printf("Order processed! Transaction: %s\n", txnID)
	return nil
}

func main() {
	// In PRODUCTION: use real processor
	fmt.Println("=== Production Mode ===")
	realProcessor := &StripeProcessor{APIKey: "sk_live_xxx"}
	ProcessOrder(realProcessor, 99.99)

	fmt.Println()

	// In TESTING: use mock processor
	fmt.Println("=== Test Mode ===")
	mockProcessor := &MockProcessor{ShouldFail: false}
	ProcessOrder(mockProcessor, 99.99)
	fmt.Printf("Charge was called %d times\n", mockProcessor.ChargeCount)

	fmt.Println()

	// Test failure scenario
	fmt.Println("=== Test Failure Scenario ===")
	failingMock := &MockProcessor{ShouldFail: true}
	err := ProcessOrder(failingMock, 50.00)
	if err != nil {
		fmt.Println("Expected error:", err)
	}
}

