package api

import (
	"github.com/zeroberto/go-ms-template/model"
)

// ExampleAPI contains the api methods available for the Example model
type ExampleAPI interface {
	// Get provides all Examples
	Get() ([]*model.Example, error)
	// GetByID provides an Example via an ID
	GetByID(ID uint) (*model.Example, error)
	// Create creates a new Example
	Create(model model.Example) (uint, error)
	// Update updates or creates, if it does not exist, a complete Example
	Update(ID uint, model model.Example) error
	// PartialUpdate updates the properties of an existing Example
	PartialUpdate(ID uint, properties interface{}) error
	// Delete deletes an existing Example
	Delete(ID uint) error
}
