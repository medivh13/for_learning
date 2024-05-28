package usecase

import (
	bookUC "for_learning/src/app/usecase/books"
	pickUpUC "for_learning/src/app/usecase/pickup"
	bookgRPCUC "for_learning/src/app/usecase/books_grpc"
)

type AllUseCases struct {
	BookUC bookUC.BooksUCInterface
	PickUpUC pickUpUC.PickUpUCInterface
	BookGrpcUC bookgRPCUC.BooksGRPCUCInterface
}
