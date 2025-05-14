package main

import (
	"design-patterns-examples/command"
	"design-patterns-examples/composite"
	"design-patterns-examples/decorator"
	"design-patterns-examples/factory"
	"design-patterns-examples/iterator"
	"design-patterns-examples/observer"
	"design-patterns-examples/proxy"
	"design-patterns-examples/singleton"
	"design-patterns-examples/strategy"
	"fmt"
)

func main() {
	// Singleton Pattern
	fmt.Println("=================")
	fmt.Println("Singleton Pattern")
	fmt.Println("=================")
	config := singleton.GetInstance()
	fmt.Printf("Config: %v\n", config.AppConfig)
	fmt.Println("Concurrency Test:")
	singleton.TestConcurrency()
	fmt.Println()

	// Factory Pattern
	fmt.Println("=================")
	fmt.Println("Factory Pattern")
	fmt.Println("=================")
	stripe := factory.NewPaymentClient(&factory.StripeFactory{})
	paypal := factory.NewPaymentClient(&factory.PayPalFactory{})
	fmt.Println(stripe.Process(100.0))
	fmt.Println(paypal.Process(100.0))
	fmt.Println()

	// Observer Pattern
	fmt.Println("=================")
	fmt.Println("Observer Pattern")
	fmt.Println("=================")
	ticker := &observer.StockTicker{Symbol: "AAPL", Price: 150.0, Observers: []observer.Observer{}}
	investor1 := &observer.Investor{Name: "Alice"}
	investor2 := &observer.Investor{Name: "Bob"}
	ticker.Attach(investor1)
	ticker.Attach(investor2)
	ticker.Notify("AAPL price updated to $150.0")
	fmt.Println()

	// Decorator Pattern
	fmt.Println("=================")
	fmt.Println("Decorator Pattern")
	fmt.Println("=================")
	handler := decorator.CompositeHandler()
	fmt.Println(handler("authenticated"))
	fmt.Println(handler("unauthenticated"))
	fmt.Println()

	// Strategy Pattern
	fmt.Println("=================")
	fmt.Println("Strategy Pattern")
	fmt.Println("=================")
	checkout := &strategy.CheckoutContext{}
	checkout.SetStrategy(&strategy.PercentDiscount{Rate: 0.1})
	fmt.Printf("Percent Discount (10%% off $100): $%.2f\n", checkout.CalculateFinalAmount(100.0))
	checkout.SetStrategy(&strategy.FixedDiscount{Offset: 20.0})
	fmt.Printf("Fixed Discount ($20 off $100): $%.2f\n", checkout.CalculateFinalAmount(100.0))
	fmt.Println()

	// Proxy Pattern
	fmt.Println("=================")
	fmt.Println("Proxy Pattern")
	fmt.Println("=================")
	proxy := proxy.NewCacheProxy("123", "Sample Data")
	fmt.Println("First fetch:", proxy.FetchData())
	fmt.Println("Second fetch (cached):", proxy.FetchData())
	fmt.Println()

	// Command Pattern
	fmt.Println("=================")
	fmt.Println("Command Pattern")
	fmt.Println("=================")
	db := &command.DatabaseReceiver{}
	cmd := &command.InsertCommand{
		Receiver:   db,
		Table:      "users",
		Columns:    []string{"name", "email"},
		Values:     []string{"John", "john@example.com"},
		PrevValues: make(map[string]string),
	}
	cmd.Execute()
	cmd.Undo()
	fmt.Println()

	// Composite Pattern
	fmt.Println("=================")
	fmt.Println("Composite Pattern")
	fmt.Println("=================")
	root := &composite.Directory{Name: "root"}
	home := &composite.Directory{Name: "home"}
	file1 := &composite.File{Name: "file1.txt", Size: 100}
	file2 := &composite.File{Name: "file2.txt", Size: 200}
	root.Add(home)
	home.Add(file1)
	home.Add(file2)
	fmt.Println(root.List())
	fmt.Println()

	// Iterator Pattern
	fmt.Println("=================")
	fmt.Println("Iterator Pattern")
	fmt.Println("=================")
	collection := &iterator.StringCollection{}
	collection.AddItem("Apple")
	collection.AddItem("Banana")
	collection.AddItem("Orange")
	iterator.TraverseCollection(collection)
}
