package api

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/zeroberto/go-ms-template/chrono"
	"github.com/zeroberto/go-ms-template/chrono/provider"

	"github.com/zeroberto/go-ms-template/api"
	"github.com/zeroberto/go-ms-template/api/rest"
	"github.com/zeroberto/go-ms-template/usecase"

	"github.com/zeroberto/go-ms-template/model"
)

func TestCreate(t *testing.T) {
	expected := api.Response{
		Code: 201,
		Path: 1,
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	createExampleMock = func(example *model.Example) (*model.Example, error) {
		return example, nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
	}
	got := eapi.Create(model.Example{ID: 1})

	if expected != got {
		t.Errorf("Create() failed, expected %v, got %v", expected, got)
	}
}

func TestCreateWhenECUCReturnsErrorThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 400,
		Body: api.ResponseBody{
			Code:    400,
			Message: "error",
			Time:    currentTime,
		},
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	createExampleMock = func(example *model.Example) (*model.Example, error) {
		return nil, &usecase.Error{Cause: errors.New("error")}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
		TS:   &timeStampMock{},
	}
	got := eapi.Create(model.Example{ID: 1})

	if expected != got {
		t.Errorf("Create() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdate(t *testing.T) {
	expected := api.Response{
		Code: 204,
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	updateExampleMock = func(example *model.Example) (*model.Example, error) {
		return example, nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
	}
	got := eapi.Update(1, model.Example{})

	if expected != got {
		t.Errorf("Update() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateWhenNotExistsThenCreateSuccess(t *testing.T) {
	expected := api.Response{
		Code: 201,
		Path: 1,
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	updateExampleMock = func(example *model.Example) (*model.Example, error) {
		return nil, &usecase.NotExistsError{}
	}
	createExampleMock = func(example *model.Example) (*model.Example, error) {
		example.ID = 1
		return example, nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
	}
	got := eapi.Update(1, model.Example{})

	if expected != got {
		t.Errorf("Update() failed, expected %v, got %v", expected, got)
	}
}

func TestUpdateWhenECUCReturnsErrorThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 400,
		Body: api.ResponseBody{
			Code:    400,
			Message: "error",
			Time:    currentTime,
		},
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	updateExampleMock = func(example *model.Example) (*model.Example, error) {
		return nil, &usecase.Error{Cause: errors.New("error")}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
		TS:   &timeStampMock{},
	}
	got := eapi.Update(1, model.Example{})

	if expected != got {
		t.Errorf("Update() failed, expected %v, got %v", expected, got)
	}
}

func TestPartialUpdate(t *testing.T) {
	expected := api.Response{
		Code: 204,
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	updateExamplePropertiesMock = func(ID int64, properties map[string]interface{}) (*model.Example, error) {
		return &model.Example{}, nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
	}
	got := eapi.PartialUpdate(1, map[string]interface{}{
		"Name": "test",
	})

	if expected != got {
		t.Errorf("PartialUpdate() failed, expected %v, got %v", expected, got)
	}
}

func TestPartialUpdateWhenIDNotExistsThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 404,
		Body: api.ResponseBody{
			Code:    404,
			Message: "No examples found for ID 1",
			Time:    currentTime,
		},
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	updateExamplePropertiesMock = func(ID int64, properties map[string]interface{}) (*model.Example, error) {
		return nil, &usecase.NotExistsError{ID: ID}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
		TS:   &timeStampMock{},
	}
	got := eapi.PartialUpdate(1, map[string]interface{}{
		"Name": "test",
	})

	if expected != got {
		t.Errorf("PartialUpdate() failed, expected %v, got %v", expected, got)
	}
}

func TestPartialUpdateWhenECUCReturnsErrorThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 400,
		Body: api.ResponseBody{
			Code:    400,
			Message: "error",
			Time:    currentTime,
		},
	}

	var ecuc usecase.ExampleCreationUseCase = &exampleCreationUseCaseMock{}
	updateExamplePropertiesMock = func(ID int64, properties map[string]interface{}) (*model.Example, error) {
		return nil, &usecase.Error{Cause: errors.New("error")}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ECUC: ecuc,
		TS:   &timeStampMock{},
	}
	got := eapi.PartialUpdate(1, map[string]interface{}{
		"Name": "test",
	})

	if expected != got {
		t.Errorf("PartialUpdate() failed, expected %v, got %v", expected, got)
	}
}

func TestGet(t *testing.T) {
	examples := []model.Example{
		model.Example{ID: 1},
		model.Example{ID: 2},
	}
	expected := api.Response{
		Code: 200,
		Body: examples,
	}

	var eruc usecase.ExampleReadUseCase = &exampleReadUseCaseMock{}
	listExamplesMock = func() ([]model.Example, error) {
		return examples, nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERUC: eruc,
	}
	got := eapi.Get()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Get() failed, expected %v, got %v", expected, got)
	}
}

func TestGetWhenERUCReturnsErrorThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 400,
		Body: api.ResponseBody{
			Code:    400,
			Message: "error",
			Time:    currentTime,
		},
	}

	var eruc usecase.ExampleReadUseCase = &exampleReadUseCaseMock{}
	listExamplesMock = func() ([]model.Example, error) {
		return nil, &usecase.Error{Cause: errors.New("error")}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERUC: eruc,
		TS:   &timeStampMock{},
	}
	got := eapi.Get()

	if expected != got {
		t.Errorf("Get() failed, expected %v, got %v", expected, got)
	}
}

func TestGetByID(t *testing.T) {
	expected := api.Response{
		Code: 200,
		Body: model.Example{ID: 1},
	}

	var eruc usecase.ExampleReadUseCase = &exampleReadUseCaseMock{}
	getExampleMock = func(ID int64) (*model.Example, error) {
		return &model.Example{ID: 1}, nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERUC: eruc,
	}
	got := eapi.GetByID(1)

	if expected != got {
		t.Errorf("GetByID() failed, expected %v, got %v", expected, got)
	}
}

func TestGetByIDWhenERUCReturnsErrorThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 400,
		Body: api.ResponseBody{
			Code:    400,
			Message: "error",
			Time:    currentTime,
		},
	}

	var eruc usecase.ExampleReadUseCase = &exampleReadUseCaseMock{}
	getExampleMock = func(ID int64) (*model.Example, error) {
		return nil, &usecase.Error{Cause: errors.New("error")}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERUC: eruc,
		TS:   &timeStampMock{},
	}
	got := eapi.GetByID(1)

	if expected != got {
		t.Errorf("GetByID() failed, expected %v, got %v", expected, got)
	}
}

func TestDelete(t *testing.T) {
	expected := api.Response{
		Code: 204,
	}

	var ermuc usecase.ExampleRemovalUseCase = &exampleRemovalUseCaseMock{}
	deleteExampleMock = func(ID int64) error {
		return nil
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERMUC: ermuc,
	}
	got := eapi.Delete(1)

	if expected != got {
		t.Errorf("Delete() failed, expected %v, got %v", expected, got)
	}
}

func TestDeleteWhenIDNotExistsThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 404,
		Body: api.ResponseBody{
			Code:    404,
			Message: "No examples found for ID 1",
			Time:    currentTime,
		},
	}

	var ermuc usecase.ExampleRemovalUseCase = &exampleRemovalUseCaseMock{}
	deleteExampleMock = func(ID int64) error {
		return &usecase.NotExistsError{ID: ID}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERMUC: ermuc,
		TS:    &timeStampMock{},
	}
	got := eapi.Delete(1)

	if expected != got {
		t.Errorf("Delete() failed, expected %v, got %v", expected, got)
	}
}

func TestDeleteWhenERMUCReturnsErrorThenFailure(t *testing.T) {
	expected := api.Response{
		Code: 400,
		Body: api.ResponseBody{
			Code:    400,
			Message: "error",
			Time:    currentTime,
		},
	}

	var ermuc usecase.ExampleRemovalUseCase = &exampleRemovalUseCaseMock{}
	deleteExampleMock = func(ID int64) error {
		return &usecase.Error{Cause: errors.New("error")}
	}

	var eapi api.ExampleAPI = &rest.ExampleAPIRest{
		ERMUC: ermuc,
		TS:    &timeStampMock{},
	}
	got := eapi.Delete(1)

	if expected != got {
		t.Errorf("Delete() failed, expected %v, got %v", expected, got)
	}
}

var currentTime time.Time = time.Now()

var timeStamp chrono.TimeStamp = &provider.TimeStampImpl{}

var createExampleMock func(example *model.Example) (*model.Example, error)

var updateExampleMock func(example *model.Example) (*model.Example, error)

var updateExamplePropertiesMock func(ID int64, properties map[string]interface{}) (*model.Example, error)

var listExamplesMock func() ([]model.Example, error)

var getExampleMock func(ID int64) (*model.Example, error)

var deleteExampleMock func(ID int64) error

type exampleCreationUseCaseMock struct{}

type exampleReadUseCaseMock struct{}

type exampleRemovalUseCaseMock struct{}

type timeStampMock struct{}

func (ecuc *exampleCreationUseCaseMock) CreateExample(example *model.Example) (*model.Example, error) {
	return createExampleMock(example)
}

func (ecuc *exampleCreationUseCaseMock) UpdateExample(example *model.Example) (*model.Example, error) {
	return updateExampleMock(example)
}

func (ecuc *exampleCreationUseCaseMock) UpdateExampleProperties(ID int64, properties map[string]interface{}) (*model.Example, error) {
	return updateExamplePropertiesMock(ID, properties)
}

func (eruc *exampleReadUseCaseMock) ListExamples() ([]model.Example, error) {
	return listExamplesMock()
}

func (eruc *exampleReadUseCaseMock) ListActiveExamples() ([]model.Example, error) {
	return nil, nil
}

func (eruc *exampleReadUseCaseMock) GetExample(ID int64) (*model.Example, error) {
	return getExampleMock(ID)
}

func (eruc *exampleReadUseCaseMock) GetExampleByName(name string) (*model.Example, error) {
	return nil, nil
}

func (ermuc *exampleRemovalUseCaseMock) DeleteExample(ID int64) error {
	return deleteExampleMock(ID)
}

func (ermuc *exampleRemovalUseCaseMock) DeleteExampleLogically(ID int64, deactivationDatetime time.Time) error {
	return nil
}

func (tp *timeStampMock) GetCurrentTime() time.Time {
	return currentTime
}
