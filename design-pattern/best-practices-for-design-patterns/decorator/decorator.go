package decorator

import "fmt"

// Handler Basic processing function
type Handler func(request string) string

// LogDecorator Log decorator
func LogDecorator(next Handler) Handler {
	return func(request string) string {
		fmt.Printf("Received request: %s\n", request)
		return next(request)
	}
}

// AuthDecorator Authentication decorator
func AuthDecorator(next Handler) Handler {
	return func(request string) string {
		if request != "authenticated" {
			return "Unauthorized"
		}
		return next(request)
	}
}

// Basic processing function
func BasicHandler(request string) string {
	return fmt.Sprintf("Processed: %s", request)
}

// Composite decorator example
func CompositeHandler() Handler {
	return LogDecorator(AuthDecorator(BasicHandler))
}
