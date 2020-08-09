package rest

import (
	"github.com/zeroberto/go-ms-template/model"
)

// ExampleAPIRest is responsible for implementing the ExampleAPIInterface using HTTP REST abstraction
type ExampleAPIRest struct {
}

// Get provides all Examples by REST abstraction
func (exampleApi *ExampleAPIRest) Get() (*model.Example, error) {
	return nil, nil
}
