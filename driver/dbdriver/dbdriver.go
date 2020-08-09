package dbdriver

import (
	"database/sql"

	"github.com/pkg/errors"
)

// SQLDriver is responsible for performing operations on a SQL database
type SQLDriver interface {
	// Execute an SQL statement with transaction
	Execute(query string, args ...interface{}) (sql.Result, error)
	// Prepare a sql statement with transaction for future execution
	PrepareAndExecute(query string, args ...interface{}) (sql.Result, error)
	// Query is responsible for executing an sql command and returning multiple lines
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryRow is responsible for executing an sql command and returning a single line
	QueryRow(query string, args ...interface{}) *sql.Row
	// Transactional is used to indicate that the current interface is enabled to use transactions
	Transactional
}

// NoSQLDriver is responsible for performing operations on a NoSQL database
type NoSQLDriver interface {
	// AddDoc is responsible for creating a new document in the database
	AddDoc(doc interface{}, args ...interface{}) (result *interface{}, err error)
	// ReplaceDoc is responsible for replacing an existing document in the database
	ReplaceDoc(UID string, doc interface{}, args ...interface{}) (result *interface{}, err error)
	// Command is responsible for executing an command in the database
	Command(command string, args ...interface{}) (result *interface{}, err error)
	// QueryDocs is responsible for obtaining several documents from the database
	QueryDocs(query string, args ...interface{}) ([]interface{}, error)
	// QueryDoc is responsible for obtaining a single document from the database
	QueryDoc(query string, args ...interface{}) (result *interface{}, err error)
}

// GraphDriver is responsible for performing operations on graph databases
type GraphDriver interface {
	// AddElement is responsible for creating a new element on the graph database
	AddElement(element interface{}, elementType interface{}, args ...interface{}) (result *interface{}, err error)
	// Command is responsible for executing an command in the graph database
	Command(command string, args ...interface{}) (result *interface{}, err error)
	// QueryElements is responsible for obtaining several elements from the graph database
	QueryElements(query string, args ...interface{}) ([]interface{}, error)
	// QueryElement is responsible for obtaining a single element from the graph database
	QueryElement(query string, args ...interface{}) (result *interface{}, err error)
}

// Transactional is responsible for enabling transactions
type Transactional interface {
	// BeginTransaction is responsible for initiating a transaction in the database
	BeginTransaction() error
	// Commit is responsible for persisting the modified information in the database
	Commit() error
	// EndTransaction is responsible for ending the transaction in the database
	EndTransaction() error
	// Rollback is responsible for undoing all modifications made to the database within the current transaction
	Rollback() error
}

// DatabaseDriverError is responsible for encapsulating errors generated by operations in the database
type DatabaseDriverError struct {
	Cause   error
	Message string
}

func (err *DatabaseDriverError) Error() string {
	return errors.Wrap(err.Cause, err.Message).Error()
}