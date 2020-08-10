package removal

import (
	"time"

	"github.com/zeroberto/go-ms-template/dataservice"
	"github.com/zeroberto/go-ms-template/usecase"
)

// ExampleRemovalUseCaseImpl corresponds to the implementation of the example model removal use case
type ExampleRemovalUseCaseImpl struct {
	EDS dataservice.ExampleDataService
}

// DeleteExample is responsible for permanently removing an Example model
func (eruc *ExampleRemovalUseCaseImpl) DeleteExample(ID int64) error {
	if err := eruc.notExistsByID(ID); err != nil {
		return err
	}
	err := eruc.EDS.Delete(ID)
	if err != nil {
		return &usecase.Error{Cause: err}
	}
	return nil
}

// DeleteExampleLogically is responsible for removing the Example model logically (deactivation)
func (eruc *ExampleRemovalUseCaseImpl) DeleteExampleLogically(ID int64, deactivationDatetime time.Time) error {
	if err := eruc.notExistsByID(ID); err != nil {
		return err
	}
	err := eruc.EDS.LogicalDeletion(ID, deactivationDatetime)
	if err != nil {
		return &usecase.Error{Cause: err}
	}
	return nil
}

func (eruc *ExampleRemovalUseCaseImpl) notExistsByID(ID int64) error {
	example, err := eruc.EDS.FindByID(ID)
	if err != nil {
		return &usecase.Error{Cause: err}
	}
	if example == nil {
		return &usecase.NotExistsError{ID: ID}
	}
	return nil
}
