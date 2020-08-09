package usecase

import (
	"errors"
	"testing"
	"time"

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

	dse := exampleDataServiceMock{}
	dseFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	dseCreateMock = func(example *model.Example) (persistedExample *model.Example, err error) {
		example.ID = expected.ID
		return example, nil
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

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

	dse := exampleDataServiceMock{}
	dseFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	dseCreateMock = func(example *model.Example) (persistedExample *model.Example, err error) {
		return nil, expected
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

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

	dse := exampleDataServiceMock{}
	dseFindByNameMock = func(name string) (*model.Example, error) {
		return nil, expected
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

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

	dse := exampleDataServiceMock{}
	dseFindByIDMock = func(ID int64) (*model.Example, error) {
		return &existing, nil
	}
	dseFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	dseUpdateMock = func(example *model.Example) (updatedExample *model.Example, err error) {
		return &expected, nil
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

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

	if createdAtTime != expected.CreatedAt {
		t.Errorf("UpdateExample() failed, expected %v, got %v", createdAtTime, expected.CreatedAt)
	}
}

func TestUpdateExampleDSEUpdateThenFailure(t *testing.T) {
	expected := &usecase.Error{Cause: errors.New("error")}

	dse := exampleDataServiceMock{}
	dseFindByIDMock = func(ID int64) (*model.Example, error) {
		return &model.Example{}, nil
	}
	dseFindByNameMock = func(name string) (*model.Example, error) {
		return nil, nil
	}
	dseUpdateMock = func(example *model.Example) (updatedExample *model.Example, err error) {
		return nil, expected
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

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

	dse := exampleDataServiceMock{}
	dseFindByIDMock = func(ID int64) (*model.Example, error) {
		return nil, nil
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

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

	dse := exampleDataServiceMock{}
	dseFindByIDMock = func(ID int64) (*model.Example, error) {
		return &model.Example{}, nil
	}
	dseFindByNameMock = func(name string) (*model.Example, error) {
		return &model.Example{}, nil
	}

	ecuc := creation.ExampleCreationUseCaseImpl{EDS: &dse}

	example, got := ecuc.UpdateExample(&model.Example{})

	if example != nil {
		t.Errorf("UpdateExample() failed, expected %v, got %v", nil, example)
	}

	if got == nil || expected.Error() != got.Error() {
		t.Errorf("UpdateExample() failed, expected %v, got %v", expected, got)
	}
}

var dseCreateMock func(example *model.Example) (persistedExample *model.Example, err error)

var dseDeleteMock func(ID int64) error

var dseFindAllMock func() ([]model.Example, error)

var dseFindByIDMock func(ID int64) (*model.Example, error)

var dseFindByNameMock func(name string) (*model.Example, error)

var dseLogicalDeletionMock func(ID int64, deactivationDatetime time.Time) error

var dseUpdateMock func(example *model.Example) (updatedExample *model.Example, err error)

var dseUpdatePropertyMock func(propertyName string, propertyValue interface{}) error

type exampleDataServiceMock struct{}

func (eds *exampleDataServiceMock) Create(example *model.Example) (persistedExample *model.Example, err error) {
	return dseCreateMock(example)
}

func (eds *exampleDataServiceMock) Delete(ID int64) error {
	return dseDeleteMock(ID)
}

func (eds *exampleDataServiceMock) FindAll() ([]model.Example, error) {
	return dseFindAllMock()
}

func (eds *exampleDataServiceMock) FindByID(ID int64) (*model.Example, error) {
	return dseFindByIDMock(ID)
}

func (eds *exampleDataServiceMock) FindByName(name string) (*model.Example, error) {
	return dseFindByNameMock(name)
}

func (eds *exampleDataServiceMock) LogicalDeletion(ID int64, deactivationDatetime time.Time) error {
	return dseLogicalDeletionMock(ID, deactivationDatetime)
}

func (eds *exampleDataServiceMock) Update(example *model.Example) (updatedExample *model.Example, err error) {
	return dseUpdateMock(example)
}

func (eds *exampleDataServiceMock) UpdateProperty(propertyName string, propertyValue interface{}) error {
	return dseUpdatePropertyMock(propertyName, propertyValue)
}
