package usecase

import (
	bookUC "for_learning/src/app/usecase/books"
	pickUpUC "for_learning/src/app/usecase/pickup"
)

type AllUseCases struct {
	BookUC bookUC.BooksUCInterface
	PickUpUC pickUpUC.PickUpUCInterface
}
