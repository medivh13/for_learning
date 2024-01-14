package email

import (
	dto "for_learning/src/app/dto/books"

	"github.com/stretchr/testify/mock"
)

type MockBooksDTO struct {
	mock.Mock
}

func NewMockBooksDTO() *MockBooksDTO {
	return &MockBooksDTO{}
}

var _ dto.BookDTOInterface = &MockBooksDTO{}

func (m *MockBooksDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
