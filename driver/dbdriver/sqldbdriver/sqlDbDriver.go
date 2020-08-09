package sqldbdriver

import (
	"database/sql"

	"github.com/zeroberto/go-ms-template/driver/dbdriver"
)

// SQLDBDriver is responsible for performing operations on a SQL database
type SQLDBDriver struct {
	DB *sql.Tx
}

// Execute an SQL statement with transaction on the SQL database
func (driver *SQLDBDriver) Execute(query string, args ...interface{}) (sql.Result, error) {
	return driver.DB.Exec(query, args...)
}

// PrepareAndExecute a sql statement with transaction for future execution
// for the SQL database
func (driver *SQLDBDriver) PrepareAndExecute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := driver.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return nil, &dbdriver.Error{Cause: err}
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, &dbdriver.Error{Cause: err}
	}

	return result, err
}

// Query is responsible for executing an sql command and returning multiple lines
// for the SQL database
func (driver *SQLDBDriver) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return driver.DB.Query(query, args...)
}

// QueryRow is responsible for executing an sql command and returning a single line
// for the SQL database
func (driver *SQLDBDriver) QueryRow(query string, args ...interface{}) *sql.Row {
	return driver.DB.QueryRow(query, args...)
}
