package books

import (
	dto "for_learning/src/app/dto/books"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type openLibraryService struct {
}

type OpenLibraryServices interface {
	GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error)
}

func NewIntegOpenLibrary() OpenLibraryServices {
	return &openLibraryService{}
}

func (s *openLibraryService) GetBooksBySubject(subject string) (*dto.GetBooksRespDTO, error) {
	var response dto.GetBooksRespDTO

	url := "http://openlibrary.org/subjects/%s.json?"

	url = fmt.Sprintf(url, subject)
	fmt.Println(url, "here")

	getReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request to Open Library : %v", err)
	}

	getReq.Header["Accept"] = []string{"application/json"}
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(getReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request to Open Library: %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println(err)
		bodyBytes, err := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		return nil, fmt.Errorf("failed to decode response: %v [RESP BODY: %s]", err, bodyString)
	}

	return &response, nil
}
