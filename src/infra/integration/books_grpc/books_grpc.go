package books

import (
	"context"
	"fmt"
	dto "for_learning/src/app/dto/books"
	booksProto "for_learning/src/app/proto/books"

	"github.com/sony/gobreaker"
)

type openLibraryService struct {
	circuitBreaker *gobreaker.CircuitBreaker
	client         booksProto.BookServiceClient
}

type OpenLibraryServices interface {
	GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error)
}

func NewIntegOpenLibrary(c *gobreaker.CircuitBreaker, client booksProto.BookServiceClient) OpenLibraryServices {
	return &openLibraryService{
		circuitBreaker: c,
		client:         client,
	}
}

func (s *openLibraryService) GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error) {
	var response dto.GetBooksRespDTO

	result, err := s.circuitBreaker.Execute(func() (interface{}, error) {
		req := &booksProto.BookReq{Subject: subject}
		resp, err := s.client.Book(context.Background(), req)
		if err != nil {
			return nil, fmt.Errorf("failed to get books from gRPC server: %v", err)
		}

		// Map the response to your DTO
		response = dto.GetBooksRespDTO{
			Name:        resp.Name,
			SubjectType: resp.SubjectType,
			Works:       make([]*dto.WorkDTO, len(resp.Works)),
		}
		for i, work := range resp.Works {
			response.Works[i] = &dto.WorkDTO{
				Title:        work.Title,
				CoverID:      work.CoverId,
				EditionCount: work.EditionCount,
				Authors:      make([]*dto.AuthorDTO, len(work.Authors)),
			}
			for j, author := range work.Authors {
				response.Works[i].Authors[j] = &dto.AuthorDTO{
					Name: author.Name,
				}
			}
		}

		return &response, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*dto.GetBooksRespDTO), nil
}
