package datamysql

import "github.com/zeroberto/go-ms-template/model"

// ExampleDataServiceCouchbase is responsible for providing the methods of accessing
// the data of the Example model in a Couchbase Database
type ExampleDataServiceCouchbase struct {
}

// Delete is responsible for physically removing Example from the repository
// in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) Delete(ID uint) error {
	return nil
}

// FindAll is responsible for returning all examples from the repository
// in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) FindAll() ([]model.Example, error) {
	return nil, nil
}

// FindByID is responsible for returning an Example from the repository
// in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) FindByID(ID uint) (*model.Example, error) {
	return nil, nil
}

// LogicalDeletion is responsible for removing Example logically from the repository
// in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) LogicalDeletion(ID uint) error {
	return nil
}

// Persist is responsible for persisting an Example in the repository
// in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) Persist(example *model.Example) (persistedExample *model.Example, err error) {
	return nil, nil
}

// Update is responsible for updating an existing Example
// in the repository in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) Update(example *model.Example) (updatedExample *model.Example, err error) {
	return nil, nil
}

// UpdateProperty is responsible for updating a particular Example property in the repository
// in a Couchbase Database
func (ds *ExampleDataServiceCouchbase) UpdateProperty(propertyName string, propertyValue interface{}) error {
	return nil
}
