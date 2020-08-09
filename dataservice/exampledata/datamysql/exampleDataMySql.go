package datamysql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeroberto/go-ms-template/dataservice"
	"github.com/zeroberto/go-ms-template/driver/dbdriver/sqldbdriver"
	"github.com/zeroberto/go-ms-template/model"
)

const (
	// DeleteExample represents a sql command to physically remove an Example from the base
	DeleteExample string = `DELETE FROM example WHERE id = ?`
	// PersistExample represents a sql command to insert an Example into the base
	PersistExample string = `INSERT INTO example (name, useful, created_at) VALUES (?, ?, ?)`
	// QueryExample represents a search query for Examples in the base
	QueryExample string = `SELECT * FROM example`
	// QueryExampleByID represents a search query for Example by ID in the base
	QueryExampleByID string = `SELECT * FROM example WHERE id = ?`
	// QueryExampleByName represents a search query for Example by name in the base
	QueryExampleByName string = `SELECT * FROM example WHERE id = ?`
	// UpdateExample represents a sql command to update an Example in the base
	UpdateExample string = `UPDATE example SET name = ?, useful = ? WHERE id = ?`
	// UpdateExampleProperties represents a sql command to update an Example in the base
	UpdateExampleProperties string = `UPDATE example SET %s WHERE id = ?`
	// DeactivateExample represents a sql command to update the deactivate column of the Example in the base
	DeactivateExample string = `UPDATE example SET deactivated_at = ? WHERE id = ?`
)

// ExampleDataServiceMySQL is responsible for providing the methods of accessing
// the data of the Example model in a MySQL Database
type ExampleDataServiceMySQL struct {
	sqlDriver sqldbdriver.SQLDBDriver
}

// Create is responsible for persisting an Example in the repository
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) Create(example *model.Example) (persistedExample *model.Example, err error) {
	rows, err := ds.sqlDriver.PrepareAndExecute(
		DeleteExample,
		example.Name,
		example.Useful,
		example.CreatedAt,
	)
	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}

	lastInsertID, err := rows.LastInsertId()
	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}

	example.ID = lastInsertID

	return example, nil
}

// Delete is responsible for physically removing Example from the repository
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) Delete(ID int64) error {
	_, err := ds.sqlDriver.Execute(DeleteExample, ID)

	if err != nil {
		return &dataservice.Error{Cause: err}
	}

	return nil
}

// FindAll is responsible for returning all examples from the repository
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) FindAll() ([]model.Example, error) {
	rows, err := ds.sqlDriver.Query(QueryExample)

	defer rows.Close()

	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}

	examples := []model.Example{}

	for rows.Next() {
		example, err := rowsToExample(rows)
		if err != nil {
			return nil, &dataservice.Error{Cause: err}
		}
		examples = append(examples, *example)
	}

	return examples, nil
}

// FindByID is responsible for returning an Example from the repository
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) FindByID(ID int64) (*model.Example, error) {
	rows, err := ds.sqlDriver.Query(QueryExampleByID)

	defer rows.Close()

	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}

	return toExample(rows)
}

// FindByName is responsible for returning an Example from the repository according to the name
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) FindByName(name string) (*model.Example, error) {
	rows, err := ds.sqlDriver.Query(QueryExampleByName, name)

	defer rows.Close()

	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}

	return toExample(rows)
}

// LogicalDeletion is responsible for removing Example logically from the repository
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) LogicalDeletion(ID int64, deactivationDatetime time.Time) error {
	_, err := ds.sqlDriver.PrepareAndExecute(DeactivateExample, deactivationDatetime, ID)
	if err != nil {
		return &dataservice.Error{Cause: err}
	}
	return nil
}

// Update is responsible for updating an existing Example
// in the repository in a MySQL Database
func (ds *ExampleDataServiceMySQL) Update(example *model.Example) (updatedExample *model.Example, err error) {
	rows, err := ds.sqlDriver.PrepareAndExecute(
		UpdateExample,
		example.Name,
		example.Useful,
		example.ID,
	)
	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}

	affectedRows, err := rows.RowsAffected()
	if err != nil {
		return nil, &dataservice.Error{Cause: err}
	}
	if affectedRows == 0 {
		return nil, &dataservice.Error{Message: "Couldn't update example. SQL command did not return any affected lines."}
	}

	return example, nil
}

// UpdateProperties is responsible for updating a particular Example property in the repository
// in a MySQL Database
func (ds *ExampleDataServiceMySQL) UpdateProperties(ID int64, properties map[string]interface{}) error {
	queryParams := ""
	queryParamValues := make([]interface{}, len(properties))

	for k, v := range properties {
		queryParams += fmt.Sprintf("%s = ?,", k)
		queryParamValues = append(queryParamValues, v)
	}
	strings.TrimSuffix(queryParams, ",")

	rows, err := ds.sqlDriver.PrepareAndExecute(
		fmt.Sprintf(UpdateExampleProperties, queryParams),
		queryParamValues,
	)
	if err != nil {
		return &dataservice.Error{Cause: err}
	}

	affectedRows, err := rows.RowsAffected()
	if err != nil {
		return &dataservice.Error{Cause: err}
	}
	if affectedRows == 0 {
		return &dataservice.Error{Message: "Couldn't update example properties. SQL command did not return any affected lines."}
	}
	return nil
}

func rowsToExample(rows *sql.Rows) (*model.Example, error) {
	var example model.Example
	if err := rows.Scan(
		&example.ID,
		&example.Name,
		&example.Useful,
		&example.CreatedAt,
		&example.DeactivatedAt,
	); err != nil {
		return nil, &dataservice.Error{Cause: err}
	}
	return &example, nil
}

func toExample(rows *sql.Rows) (*model.Example, error) {
	if rows.Next() {
		example, err := rowsToExample(rows)
		if err != nil {
			return nil, &dataservice.Error{Cause: err}
		}
		return example, nil
	}
	return nil, nil
}
