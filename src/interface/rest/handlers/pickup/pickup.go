package article

import (
	"encoding/json"
	"net/http"

	dto "for_learning/src/app/dto/pickup"
	usecases "for_learning/src/app/usecase/pickup"
	common_error "for_learning/src/infra/errors"
	"for_learning/src/interface/rest/response"
)

type BooksHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type booksHandler struct {
	response response.IResponseClient
	usecase  usecases.PickUpUCInterface
}

func NewBooksHandler(r response.IResponseClient, h usecases.PickUpUCInterface) BooksHandlerInterface {
	return &booksHandler{
		response: r,
		usecase:  h,
	}
}

func (h *booksHandler) Create(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.ReqPickupDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Create(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Adding New PickUp Schedule",
		nil,
		nil,
	)
}
