package read

import (
	"github.com/zeroberto/go-ms-template/dataservice"
	"github.com/zeroberto/go-ms-template/model"
	"github.com/zeroberto/go-ms-template/usecase"
)

// ExampleReadUseCaseImpl corresponds to the implementation of the example model read use case
type ExampleReadUseCaseImpl struct {
	EDS dataservice.ExampleDataService
}

// ListExamples is responsible for obtaining all registered Examples
func (eruc *ExampleReadUseCaseImpl) ListExamples() ([]model.Example, error) {
	examples, err := eruc.EDS.FindAll()
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return examples, nil
}

// ListActiveExamples is responsible for obtaining all active Examples
func (eruc *ExampleReadUseCaseImpl) ListActiveExamples() ([]model.Example, error) {
	examples, err := eruc.EDS.FindActives()
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return examples, nil
}

// GetExample is responsible for obtaining an Example according to the given identifier
func (eruc *ExampleReadUseCaseImpl) GetExample(ID int64) (*model.Example, error) {
	example, err := eruc.EDS.FindByID(ID)
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return example, nil
}

// GetExampleByName is responsible for obtaining an Example according to the given name
func (eruc *ExampleReadUseCaseImpl) GetExampleByName(name string) (*model.Example, error) {
	example, err := eruc.EDS.FindByName(name)
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return example, nil
}
