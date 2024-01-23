package books

import (
	"encoding/json"
	"fmt"
	dto "for_learning/src/app/dto/books"
	"io"
	"net/http"

	"github.com/sony/gobreaker"
)

type openLibraryService struct {
	circuitBreaker *gobreaker.CircuitBreaker
}

type OpenLibraryServices interface {
	GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error)
}

func NewIntegOpenLibrary(c *gobreaker.CircuitBreaker) OpenLibraryServices {
	return &openLibraryService{
		circuitBreaker: c,
	}
}

func (s *openLibraryService) GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error) {
	var response dto.GetBooksRespDTO

	result, err := s.circuitBreaker.Execute(func() (interface{}, error) {
		url := "http://openlibrary.org/subjects/%s.json?"

		url = fmt.Sprintf(url, subject)

		getReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create http request to Open Library : %v", err)
		}

		getReq.Header["Accept"] = []string{"application/json"}
		client := http.Client{}

		resp, err := client.Do(getReq)
		if err != nil {
			return nil, fmt.Errorf("failed to create http request to Open Library: %v", err)
		}

		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			bodyBytes, err := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			return nil, fmt.Errorf("failed to decode response: %v [RESP BODY: %s]", err, bodyString)
		}

		return &response, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*dto.GetBooksRespDTO), nil
}
