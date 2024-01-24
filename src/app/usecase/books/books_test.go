package books

import (
	"context"
	"encoding/json"
	"errors"
	mockDTO "for_learning/mocks/app/dto/books"
	mockInteg "for_learning/mocks/infra/integration/books"
	"time"

	mockRedis "for_learning/mocks/infra/persistence/redis"
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
	mockRedis *mockRedis.MockRedis
	dtoGet    *dto.BookReqDTO
	dtoResp   *dto.GetBooksRespDTO
}

func (suite *BooksUseCaseList) SetupTest() {
	suite.resp = new(mockReponse.MockResponse)
	suite.mockDTO = new(mockDTO.MockBooksDTO)
	suite.mockRedis = new(mockRedis.MockRedis)
	suite.mockInteg = new(mockInteg.MockInteg)
	suite.useCase = NewBooksUseCase(suite.mockInteg, suite.mockRedis)

	suite.dtoGet = &dto.BookReqDTO{
		Subject: "love",
	}

	suite.dtoResp = &dto.GetBooksRespDTO{
		Name: "test",
	}

}

func (u *BooksUseCaseList) TestGetBySubjectFromRedisSuccess() {
	ctx := context.Background()
	dataresp, _ := json.Marshal(u.dtoResp)
	u.mockRedis.Mock.On("GetData", ctx, "love").Return(string(dataresp), nil)
	_, err := u.useCase.GetBooksBySubject(ctx, u.dtoGet)
	u.Equal(nil, err)
}

func (u *BooksUseCaseList) TestGetBySubjectSuccess() {
	ctx := context.Background()
	u.mockRedis.Mock.On("GetData", ctx, "love").Return("", errors.New(mock.Anything))
	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(mock.Anything, nil)
	u.mockRedis.Mock.On("SetData", ctx, "love", mock.Anything, time.Duration(2)*time.Minute).Return(nil)
	_, err := u.useCase.GetBooksBySubject(ctx, u.dtoGet)
	u.Equal(nil, err)
}

func (u *BooksUseCaseList) TestGetBySubjectFail() {
	ctx := context.Background()
	u.mockRedis.Mock.On("GetData", ctx, "love").Return("", errors.New(mock.Anything))
	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetBooksBySubject(ctx, u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *BooksUseCaseList) TestGetBySubjectSetDataRedisFail() {
	ctx := context.Background()
	dataresp, _ := json.Marshal(u.dtoResp)
	u.mockRedis.Mock.On("GetData", ctx, "love").Return("", errors.New(mock.Anything))
	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(u.dtoResp, nil)
	u.mockRedis.Mock.On("SetData", ctx, "love", dataresp, time.Duration(2)*time.Minute).Return(errors.New(mock.Anything))
	_, err := u.useCase.GetBooksBySubject(ctx, u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(BooksUseCaseList))
}
