package creation

import (
	"fmt"

	"github.com/zeroberto/go-ms-template/dataservice"
	"github.com/zeroberto/go-ms-template/model"
	"github.com/zeroberto/go-ms-template/tool"
	"github.com/zeroberto/go-ms-template/usecase"
)

// ExampleCreationUseCaseImpl corresponds to the implementation of the example model creation use case
type ExampleCreationUseCaseImpl struct {
	EDS dataservice.ExampleDataService
}

// CreateExample is responsible for creating a new Example
func (ecuc *ExampleCreationUseCaseImpl) CreateExample(example *model.Example) (*model.Example, error) {
	if err := ecuc.existsByName(example.Name, example.ID); err != nil {
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
	if err := ecuc.notExistsByID(example.ID); err != nil {
		return nil, err
	}
	if err := ecuc.existsByName(example.Name, example.ID); err != nil {
		return nil, err
	}
	example, err := ecuc.EDS.Update(example)
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return example, nil
}

// UpdateExampleProperties is responsible for updating partial properties of the Example model
func (ecuc *ExampleCreationUseCaseImpl) UpdateExampleProperties(ID int64, properties map[string]interface{}) (*model.Example, error) {
	if err := ecuc.notExistsByID(ID); err != nil {
		return nil, err
	}
	propertyNames := getUpgradeableProperties()
	for k := range properties {
		if !tool.ContainsString(k, propertyNames) {
			return nil, &usecase.Error{Message: fmt.Sprintf("property %s does not exist or cannot be updated", k)}
		}
	}
	if tool.ContainsStringKey("Name", properties) {
		if err := ecuc.existsByName(properties["Name"].(string), ID); err != nil {
			return nil, err
		}
	}
	if err := ecuc.EDS.UpdateProperties(ID, properties); err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	example, err := ecuc.EDS.FindByID(ID)
	if err != nil {
		return nil, &usecase.Error{Cause: err}
	}
	return example, nil
}

func (ecuc *ExampleCreationUseCaseImpl) existsByName(name string, ID int64) error {
	example, err := ecuc.EDS.FindByName(name)
	if err != nil {
		return &usecase.Error{Cause: err}
	}
	if example != nil && example.ID != ID {
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

func getUpgradeableProperties() []string {
	return []string{"Name", "Useful"}
}
