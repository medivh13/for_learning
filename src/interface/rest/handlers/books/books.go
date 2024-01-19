package article

import (
	"net/http"

	dto "for_learning/src/app/dto/books"
	usecases "for_learning/src/app/usecase/books"
	common_error "for_learning/src/infra/errors"
	limiter "for_learning/src/infra/limiter"
	"for_learning/src/interface/rest/response"
)

type BooksHandlerInterface interface {
	GetBySubject(w http.ResponseWriter, r *http.Request)
}

type booksHandler struct {
	response response.IResponseClient
	usecase  usecases.BooksUCInterface
	limiter  limiter.RateLimiterInterface
}

func NewBooksHandler(r response.IResponseClient, h usecases.BooksUCInterface, l limiter.RateLimiterInterface) BooksHandlerInterface {
	return &booksHandler{
		response: r,
		usecase:  h,
		limiter:  l,
	}
}

func (h *booksHandler) GetBySubject(w http.ResponseWriter, r *http.Request) {

	if !h.limiter.Allow() {
		h.response.HttpError(w, common_error.NewError(common_error.RATE_LIMIT_EXCEEDED, nil))
		return
	}

	getDTO := dto.BookReqDTO{}

	if r.URL.Query().Get("subject") != "" {
		getDTO.Subject = r.URL.Query().Get("subject")
	}

	err := getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetBooksBySubject(r.Context(), &getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Books",
		data,
		nil,
	)
}
