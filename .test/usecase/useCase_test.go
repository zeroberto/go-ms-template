package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/zeroberto/go-ms-template/dataservice"

	"github.com/zeroberto/go-ms-template/usecase"

	"github.com/zeroberto/go-ms-template/usecase/creation"

	"github.com/zeroberto/go-ms-template/model"
)

func TestCreateExample(t *testing.T) {
	fixedTime := time.Now()

	expected := model.Example{
		ID:        1,
		Name:      "test",
		Useful:    true,
		CreatedAt: fixedTime,
	}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	edsCreateMock = func(example *model.Example) (persistedExample *model.Example, err error) {
		example.ID = expected.ID
		return example, nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	got, err := ecuc.CreateExample(&model.Example{
		Name:      "test",
		Useful:    true,
		CreatedAt: fixedTime,
	})

	if err != nil {
		t.Errorf("CreateExample() failed, error %v", err)
	}

	if expected != *got {
		t.Errorf("CreateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestCreateExampleWhenDSECreateReturnsErrorThenFailure(t *testing.T) {
	expected := &usecase.Error{Cause: errors.New("error")}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	edsCreateMock = func(example *model.Example) (persistedExample *model.Example, err error) {
		return nil, expected
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.CreateExample(&model.Example{})

	if example != nil {
		t.Errorf("CreateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("CreateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestCreateExampleWhenNameAlreadyExistsThenFailure(t *testing.T) {
	expected := &usecase.Error{Message: "Example already exists"}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return nil, expected
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.CreateExample(&model.Example{})

	if example != nil {
		t.Errorf("CreateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("CreateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExample(t *testing.T) {
	currentFixedTime := time.Now()
	createdAtTime := currentFixedTime.Add(-1 * time.Minute)

	expected := model.Example{
		ID:        1,
		Name:      "updated",
		Useful:    true,
		CreatedAt: createdAtTime,
	}
	existing := model.Example{
		ID:        1,
		Name:      "existing",
		Useful:    false,
		CreatedAt: createdAtTime,
	}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return &existing, nil
	}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	edsUpdateMock = func(example *model.Example) (updatedExample *model.Example, err error) {
		return &expected, nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	got, err := ecuc.UpdateExample(&model.Example{
		Name:      "updated",
		Useful:    true,
		CreatedAt: currentFixedTime,
	})

	if err != nil {
		t.Errorf("UpdateExample() failed, error %v", err)
	}

	if expected != *got {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExampleDSEUpdateThenFailure(t *testing.T) {
	expected := &usecase.Error{Cause: errors.New("error")}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return &model.Example{}, nil
	}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	edsUpdateMock = func(example *model.Example) (updatedExample *model.Example, err error) {
		return nil, expected
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.UpdateExample(&model.Example{})

	if example != nil {
		t.Errorf("UpdateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExampleWhenIDNotExistsThenFailure(t *testing.T) {
	expected := &usecase.Error{Message: "No example found for this ID"}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return nil, nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.UpdateExample(&model.Example{})

	if example != nil {
		t.Errorf("UpdateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExampleWhenNameAlreadyExistsThenFailure(t *testing.T) {
	expected := &usecase.Error{Message: "Example already exists"}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return &model.Example{ID: 1}, nil
	}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return &model.Example{ID: 2}, nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.UpdateExample(&model.Example{ID: 1})

	if example != nil {
		t.Errorf("UpdateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExamplePropertiesWhenPropertyIsUsefulThenSuccess(t *testing.T) {
	currentFixedTime := time.Now()
	createdAtTime := currentFixedTime.Add(-1 * time.Minute)

	expected := model.Example{
		ID:        1,
		Name:      "shouldNotUpdated",
		Useful:    true,
		CreatedAt: createdAtTime,
	}
	existing := model.Example{
		ID:        1,
		Name:      "shouldNotUpdated",
		Useful:    false,
		CreatedAt: createdAtTime,
	}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return &existing, nil
	}
	edsUpdatePropertiesMock = func(ID int64, properties map[string]interface{}) error {
		existing.Useful = properties["Useful"].(bool)
		return nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	got, err := ecuc.UpdateExampleProperties(1, map[string]interface{}{
		"Useful": true,
	})

	if err != nil {
		t.Errorf("UpdateExampleProperties() failed, error %v", err)
	}

	if expected != *got {
		t.Errorf("UpdateExampleProperties() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExamplePropertiesWhenPropertyNotExistsThenFailure(t *testing.T) {
	expected := &usecase.Error{Message: "property Wrong does not exist or cannot be updated"}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return &model.Example{}, nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.UpdateExampleProperties(1, map[string]interface{}{
		"Wrong": 1,
	})

	if example != nil {
		t.Errorf("UpdateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateExamplePropertiesWhenPropertyIsNameAndNameAlreadyExistsThenFailure(t *testing.T) {
	expected := &usecase.Error{Message: "Example already exists"}

	var eds dataservice.ExampleDataService = &exampleDataServiceMock{}
	edsFindByIDMock = func(ID int64) (*model.Example, error) {
		return &model.Example{ID: 1}, nil
	}
	edsFindByNameMock = func(name string) (*model.Example, error) {
		return &model.Example{ID: 2}, nil
	}

	var ecuc usecase.ExampleCreationUseCase = &creation.ExampleCreationUseCaseImpl{EDS: eds}

	example, got := ecuc.UpdateExampleProperties(1, map[string]interface{}{
		"Name": "shouldNotUpdated",
	})

	if example != nil {
		t.Errorf("UpdateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

var edsCreateMock func(example *model.Example) (persistedExample *model.Example, err error)

var edsDeleteMock func(ID int64) error

var edsFindAllMock func() ([]model.Example, error)

var edsFindByIDMock func(ID int64) (*model.Example, error)

var edsFindByNameMock func(name string) (*model.Example, error)

var edsLogicalDeletionMock func(ID int64, deactivationDatetime time.Time) error

var edsUpdateMock func(example *model.Example) (updatedExample *model.Example, err error)

var edsUpdatePropertiesMock func(ID int64, properties map[string]interface{}) error

type exampleDataServiceMock struct{}

func (eds *exampleDataServiceMock) Create(example *model.Example) (persistedExample *model.Example, err error) {
	return edsCreateMock(example)
}

func (eds *exampleDataServiceMock) Delete(ID int64) error {
	return edsDeleteMock(ID)
}

func (eds *exampleDataServiceMock) FindAll() ([]model.Example, error) {
	return edsFindAllMock()
}

func (eds *exampleDataServiceMock) FindByID(ID int64) (*model.Example, error) {
	return edsFindByIDMock(ID)
}

func (eds *exampleDataServiceMock) FindByName(name string) (*model.Example, error) {
	return edsFindByNameMock(name)
}

func (eds *exampleDataServiceMock) LogicalDeletion(ID int64, deactivationDatetime time.Time) error {
	return edsLogicalDeletionMock(ID, deactivationDatetime)
}

func (eds *exampleDataServiceMock) Update(example *model.Example) (updatedExample *model.Example, err error) {
	return edsUpdateMock(example)
}

func (eds *exampleDataServiceMock) UpdateProperties(ID int64, properties map[string]interface{}) error {
	return edsUpdatePropertiesMock(ID, properties)
}
