package rest

import (
	"net/http"
	"time"

	"github.com/zeroberto/go-ms-template/dataservice"

	"github.com/zeroberto/go-ms-template/api"
	"github.com/zeroberto/go-ms-template/chrono"
	"github.com/zeroberto/go-ms-template/model"
	"github.com/zeroberto/go-ms-template/usecase"
)

// ExampleAPIRest is responsible for implementing the ExampleAPIInterface using HTTP REST abstraction
type ExampleAPIRest struct {
	ECUC  usecase.ExampleCreationUseCase
	ERUC  usecase.ExampleReadUseCase
	ERMUC usecase.ExampleRemovalUseCase
	TS    chrono.TimeStamp
}

// Get provides all Examples by REST abstraction
func (eapi *ExampleAPIRest) Get() api.Response {
	examples, err := eapi.ERUC.ListExamples()
	if err != nil {
		return report(err, eapi.TS.GetCurrentTime())
	}
	return api.Response{
		Code: http.StatusOK,
		Body: examples,
	}
}

// GetByID provides an Example via an ID by REST abstraction
func (eapi *ExampleAPIRest) GetByID(ID int64) api.Response {
	example, err := eapi.ERUC.GetExample(ID)
	if err != nil {
		return report(err, eapi.TS.GetCurrentTime())
	}
	return api.Response{
		Code: http.StatusOK,
		Body: *example,
	}
}

// Create creates a new Example by REST abstraction
func (eapi *ExampleAPIRest) Create(example model.Example) api.Response {
	_, err := eapi.ECUC.CreateExample(&example)
	if err != nil {
		return report(err, eapi.TS.GetCurrentTime())
	}
	return api.Response{
		Code: http.StatusCreated,
		Path: example.ID,
	}
}

// Update updates or creates, if it does not exist, a complete Example by REST abstraction
func (eapi *ExampleAPIRest) Update(ID int64, example model.Example) api.Response {
	_, updateErr := eapi.ECUC.UpdateExample(&example)
	if updateErr != nil {
		_, ok := updateErr.(*usecase.NotExistsError)
		if ok {
			return eapi.Create(example)
		}
		return report(updateErr, eapi.TS.GetCurrentTime())
	}
	return api.Response{
		Code: http.StatusNoContent,
	}
}

// PartialUpdate updates the properties of an existing Example by REST abstraction
func (eapi *ExampleAPIRest) PartialUpdate(ID int64, properties map[string]interface{}) api.Response {
	_, err := eapi.ECUC.UpdateExampleProperties(ID, properties)
	if err != nil {
		return report(err, eapi.TS.GetCurrentTime())
	}
	return api.Response{
		Code: http.StatusNoContent,
	}
}

// Delete deletes an existing Example by REST abstraction
func (eapi *ExampleAPIRest) Delete(ID int64) api.Response {
	err := eapi.ERMUC.DeleteExample(ID)
	if err != nil {
		return report(err, eapi.TS.GetCurrentTime())
	}
	return api.Response{Code: http.StatusNoContent}
}

func report(err error, time time.Time) api.Response {
	code := getCode(err)
	return api.Response{
		Code: code,
		Body: api.ResponseBody{
			Time:    time,
			Code:    code,
			Message: err.Error(),
		},
	}
}

func getCode(err error) int {
	if _, ok := err.(*usecase.NotExistsError); ok {
		return http.StatusNotFound
	}
	if _, ok := err.(*dataservice.Error); ok {
		return http.StatusInternalServerError
	}
	if e, ok := err.(*usecase.Error); ok {
		return getCode(e.Cause)
	}
	return http.StatusBadRequest
}
