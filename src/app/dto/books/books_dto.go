package books

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BookDTOInterface interface {
	Validate() error
}

type BookReqDTO struct {
	Subject string `json:"subject"`
}

func (dto *BookReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Subject, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetBooksRespDTO struct {
	Name        string     `json:"name"`
	SubjectType string     `json:"subject_type"`
	Works       []*WorkDTO `json:"works"`
}

type WorkDTO struct {
	Title        string       `json:"title"`
	CoverID      int64        `json:"cover_id"`
	EditionCount int64        `json:"edition_count"`
	Authors      []*AuthorDTO `json:"authors"`
}

type AuthorDTO struct {
	Name string `json:"name`
}
