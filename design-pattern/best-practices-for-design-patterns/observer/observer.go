package observer

import "fmt"

// Observer Observer interface
type Observer interface {
	Update(message string)
}

// Subject Subject interface
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

// StockTicker Specific subject: stock quotation
type StockTicker struct {
	Observers []Observer
	Symbol    string
	Price     float64
}

func (s *StockTicker) Attach(observer Observer) {
	s.Observers = append(s.Observers, observer)
}

func (s *StockTicker) Detach(observer Observer) {
	for i, o := range s.Observers {
		if o == observer {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
			return
		}
	}
}

func (s *StockTicker) Notify(message string) {
	for _, o := range s.Observers {
		o.Update(message)
	}
}

// Investor Specific observer: investor
type Investor struct {
	Name string
}

func (i *Investor) Update(message string) {
	fmt.Printf("%s received: %s\n", i.Name, message)
}
