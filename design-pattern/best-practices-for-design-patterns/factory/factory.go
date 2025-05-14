package factory

import "fmt"

// PaymentProcessor Payment processor interface
type PaymentProcessor interface {
	Process(amount float64) string
}

// StripeProcessor Stripe payment implementation
type StripeProcessor struct{}

func (s *StripeProcessor) Process(amount float64) string {
	return fmt.Sprintf("Stripe processed $%.2f", amount)
}

// PayPalProcessor PayPal payment implementation
type PayPalProcessor struct{}

func (p *PayPalProcessor) Process(amount float64) string {
	return fmt.Sprintf("PayPal processed $%.2f", amount)
}

// PaymentFactory Factory interface
type PaymentFactory interface {
	CreateProcessor() PaymentProcessor
}

// StripeFactory Stripe factory implementation
type StripeFactory struct{}

func (s *StripeFactory) CreateProcessor() PaymentProcessor {
	return &StripeProcessor{}
}

// PayPalFactory PayPal factory implementation
type PayPalFactory struct{}

func (p *PayPalFactory) CreateProcessor() PaymentProcessor {
	return &PayPalProcessor{}
}

// Client usage example
func NewPaymentClient(factory PaymentFactory) PaymentProcessor {
	return factory.CreateProcessor()
}
