package response_test

import (
	"net/http"

	"for_learning/src/infra/errors"
	"for_learning/src/interface/rest/response"

	"github.com/stretchr/testify/mock"
)

type MockResponse struct {
	mock.Mock
}

func NewMockResponse() *MockResponse {
	return &MockResponse{}
}

var _ response.IResponseClient = &MockResponse{}

func (m *MockResponse) BuildMeta(page int, perPage int, count int64) *response.Meta {
	args := m.Called(page, perPage, count)
	if meta, ok := args.Get(0).(*response.Meta); ok {
		return meta
	}

	return nil
}

func (m *MockResponse) HttpError(w http.ResponseWriter, err error) error {
	args := m.Called(w, err)
	if err, ok := args.Get(0).(error); ok {
		return err
	}

	var respError errors.HttpError

	if cerr, ok := err.(*errors.CommonError); ok {
		respError = cerr.ToHttpError()
	} else {
		respError = errors.NewError(errors.UNKNOWN_ERROR, err).ToHttpError()
	}

	w.WriteHeader(respError.GetHttpStatus())
	return nil
}

func (m *MockResponse) JSON(w http.ResponseWriter, message string, data interface{}, meta *response.Meta) error {
	args := m.Called(w, message, data, meta)

	if err, ok := args.Get(0).(error); ok {
		return err
	}

	return nil
}
