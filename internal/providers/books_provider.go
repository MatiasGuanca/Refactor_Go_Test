package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"educabot.com/bookshop/internal/models"
)

type HttpBooksProvider struct {
	client  http.Client
	baseURL string
}

func NewHttpBooksProvider(baseURL string) HttpBooksProvider {
	return HttpBooksProvider{
		client:  http.Client{Timeout: 5 * time.Second},
		baseURL: baseURL,
	}
}

func (p HttpBooksProvider) GetBooks(ctx context.Context) []models.Book {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.baseURL+"/books", nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		return []models.Book{}
	}

	resp, err := p.client.Do(req)
	if err != nil {
		fmt.Println("error performing request:", err)
		return []models.Book{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("unexpected status:", resp.StatusCode)
		return []models.Book{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response:", err)
		return []models.Book{}
	}

	var books []models.Book
	if err := json.Unmarshal(body, &books); err != nil {
		fmt.Println("error decoding json:", err)
		return []models.Book{}
	}

	return books
}
