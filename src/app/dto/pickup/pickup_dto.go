package pickup

import (
	"for_learning/src/infra/errors"
	common_error "for_learning/src/infra/errors"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PickUpDTOInterface interface {
	Validate() error
}

type ReqPickupDTO struct {
	Date        string         `json:"date"`
	User        string         `json:"user"`
	Information []*Information `json:"information"`
}

type Information struct {
	Title        string       `json:"title"`
	CoverID      int64        `json:"cover_id"`
	EditionCount int64        `json:"edition_count"`
	Authors      []*AuthorDTO `json:"authors"`
}

type AuthorDTO struct {
	Name string `json:"name`
}

func (dto *ReqPickupDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Date, validation.Required, isDate("02-01-2006")),
		validation.Field(&dto.User, validation.Required),
		validation.Field(&dto.Information, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type dateFormat struct {
	format string
}

func (d dateFormat) Validate(value interface{}) error {
	dateString, ok := value.(string)
	if !ok {
		return errors.NewError(common_error.DATA_INVALID, nil)
	}

	_, err := time.Parse(d.format, dateString)
	if err != nil {
		fmt.Println(err, "here2")
		return errors.NewError(common_error.DATA_INVALID, err)
	}

	return nil
}

func isDate(format string) validation.Rule {
	return dateFormat{format}
}
