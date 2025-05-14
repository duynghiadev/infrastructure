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
	receiver   *DatabaseReceiver
	table      string
	columns    []string
	values     []string
	prevValues map[string]string // Store old values for rollback
}

func (c *InsertCommand) Execute() error {
	// Execute the insert logic and record the old values
	c.prevValues = make(map[string]string) // Initialize empty map for previous values
	return c.receiver.ExecuteSQL(fmt.Sprintf("INSERT INTO %s VALUES (%s)", c.table, c.values))
}

func (c *InsertCommand) Undo() error {
	// Use prevValues to roll back the insert operation
	return c.receiver.Rollback()
}
