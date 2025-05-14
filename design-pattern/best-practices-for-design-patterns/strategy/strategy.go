package strategy

// PromotionStrategy Promotion strategy interface
type PromotionStrategy interface {
	Calculate(amount float64) float64
}

// PercentDiscount Percentage discount strategy
type PercentDiscount struct {
	rate float64 // Discount rate (0.1 = 10%)
}

func (p *PercentDiscount) Calculate(amount float64) float64 {
	return amount * (1 - p.rate)
}

// FixedDiscount Fixed amount discount strategy
type FixedDiscount struct {
	offset float64 // Fixed reduction amount
}

func (f *FixedDiscount) Calculate(amount float64) float64 {
	if amount > f.offset {
		return amount - f.offset
	}
	return amount
}

// CheckoutContext Checkout context
type CheckoutContext struct {
	strategy PromotionStrategy
}

func (c *CheckoutContext) SetStrategy(strategy PromotionStrategy) {
	c.strategy = strategy
}

func (c *CheckoutContext) CalculateFinalAmount(amount float64) float64 {
	return c.strategy.Calculate(amount)
}
