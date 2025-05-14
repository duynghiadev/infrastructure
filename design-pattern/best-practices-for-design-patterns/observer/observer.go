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
	observers []Observer
	symbol    string
	price     float64
}

func (s *StockTicker) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *StockTicker) Detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *StockTicker) Notify(message string) {
	for _, o := range s.observers {
		o.Update(message)
	}
}

// Investor Specific observer: investor
type Investor struct {
	name string
}

func (i *Investor) Update(message string) {
	fmt.Printf("%s received: %s\n", i.name, message)
}
