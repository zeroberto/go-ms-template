package creation

import (
	"github.com/zeroberto/go-ms-template/dataservice"
	"github.com/zeroberto/go-ms-template/model"
	"github.com/zeroberto/go-ms-template/usecase"
)

// ExampleCreationUseCaseImpl corresponds to the implementation of the example model creation use case
type ExampleCreationUseCaseImpl struct {
	EDS dataservice.ExampleDataService
}

// CreateExample is responsible for creating a new Example
func (ecuc *ExampleCreationUseCaseImpl) CreateExample(example *model.Example) (*model.Example, error) {
	if err := ecuc.existsByName(example.Name); err != nil {
		return nil, err
	}
	example, err := ecuc.EDS.Create(example)
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return example, nil
}

// UpdateExample is responsible for updating the complete Example model
func (ecuc *ExampleCreationUseCaseImpl) UpdateExample(example *model.Example) (*model.Example, error) {
	if err := ecuc.existsByName(example.Name); err != nil {
		return nil, err
	}
	if err := ecuc.notExistsByID(example.ID); err != nil {
		return nil, err
	}
	example, err := ecuc.EDS.Update(example)
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return example, nil
}

// UpdateExampleProperties is responsible for updating partial properties of the Example model
func (ecuc *ExampleCreationUseCaseImpl) UpdateExampleProperties(example *model.Example) (*model.Example, error) {
	return nil, nil
}

func (ecuc *ExampleCreationUseCaseImpl) existsByName(name string) error {
	example, err := ecuc.EDS.FindByName(name)
	if err != nil {
		return &usecase.Error{Cause: err}
	}
	if example != nil {
		return &usecase.Error{Message: "Example already exists"}
	}
	return nil
}

func (ecuc *ExampleCreationUseCaseImpl) notExistsByID(ID int64) error {
	example, err := ecuc.EDS.FindByID(ID)
	if err != nil {
		return &usecase.Error{Cause: err}
	}
	if example == nil {
		return &usecase.Error{Message: "No example found for this ID"}
	}
	return nil
}
