package mock_integration

import (
	dto "for_learning/src/app/dto/books"
	integ "for_learning/src/infra/integration/books"

	"github.com/stretchr/testify/mock"
)

type MockInteg struct {
	mock.Mock
}

func NewMockInteg() *MockInteg {
	return &MockInteg{}
}

var _ integ.OpenLibraryServices = &MockInteg{}

func (o *MockInteg) GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error) {
	args := o.Called(subject)

	var (
		data *dto.GetBooksRespDTO
		err  error
	)

	if n, ok := args.Get(0).(*dto.GetBooksRespDTO); ok {
		data = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return data, err
}
