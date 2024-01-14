package books

import (
	mockDTO "for_learning/mocks/app/dto/books"
	mockInteg "for_learning/mocks/infra/integration/books"

	"errors"
	mockReponse "for_learning/mocks/interface/response"
	dto "for_learning/src/app/dto/books"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockBooksUseCase struct {
	mock.Mock
}

type BooksUseCaseList struct {
	suite.Suite
	resp    *mockReponse.MockResponse
	mockDTO *mockDTO.MockBooksDTO

	useCase   BooksUCInterface
	mockInteg *mockInteg.MockInteg

	dtoGet  *dto.BookReqDTO
	dtoResp *dto.GetBooksRespDTO
}

func (suite *BooksUseCaseList) SetupTest() {
	suite.resp = new(mockReponse.MockResponse)
	suite.mockDTO = new(mockDTO.MockBooksDTO)

	suite.mockInteg = new(mockInteg.MockInteg)
	suite.useCase = NewBooksUseCase(suite.mockInteg)

	suite.dtoGet = &dto.BookReqDTO{
		Subject: "love",
	}

	suite.dtoResp = &dto.GetBooksRespDTO{
		Name: "test",
	}

}

func (u *BooksUseCaseList) TestGetBySubjectSuccess() {

	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(mock.Anything, nil)

	_, err := u.useCase.GetBooksBySubject(u.dtoGet)
	u.Equal(nil, err)
}

func (u *BooksUseCaseList) TestGetBySubjectFail() {

	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetBooksBySubject(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(BooksUseCaseList))
}
