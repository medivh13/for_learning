package books

import (
	dto "for_learning/src/app/dto/books"
	"log"

	integ "for_learning/src/infra/integration/books"
)

type BooksUCInterface interface {
	GetBooksBySubject(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
}

type booksUseCase struct {
	BooksInteg integ.OpenLibraryServices
}

func NewBooksUseCase(i integ.OpenLibraryServices) *booksUseCase {
	return &booksUseCase{
		BooksInteg: i,
	}
}

func (uc *booksUseCase) GetBooksBySubject(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error) {

	var resp *dto.GetBooksRespDTO

	resp, err := uc.BooksInteg.GetBooksBySubject(req.Subject)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
