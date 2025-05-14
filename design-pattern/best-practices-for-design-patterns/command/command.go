package command

import (
	"database/sql"
	"fmt"
)

// DatabaseReceiver Database receiver
type DatabaseReceiver struct {
	db *sql.DB
}

func (r *DatabaseReceiver) ExecuteSQL(sqlStmt string) error {
	fmt.Printf("Executing: %s\n", sqlStmt)
	return nil // Simulate SQL execution
}

func (r *DatabaseReceiver) Rollback() error {
	fmt.Println("Rolling back transaction")
	return nil // Simulate rollback
}

// Command Command interface
type Command interface {
	Execute() error
	Undo() error
}

// InsertCommand Insert command
type InsertCommand struct {
	Receiver   *DatabaseReceiver
	Table      string
	Columns    []string
	Values     []string
	PrevValues map[string]string
}

func (c *InsertCommand) Execute() error {
	// Execute the insert logic and record the old values
	c.PrevValues = make(map[string]string) // Initialize empty map for previous values
	return c.Receiver.ExecuteSQL(fmt.Sprintf("INSERT INTO %s VALUES (%s)", c.Table, c.Values))
}

func (c *InsertCommand) Undo() error {
	// Use PrevValues to roll back the insert operation
	return c.Receiver.Rollback()
}
