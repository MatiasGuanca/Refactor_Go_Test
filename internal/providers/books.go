package providers

import (
	"context"

	"educabot.com/bookshop/internal/models"
)

type BooksProvider interface {
	GetBooks(ctx context.Context) []models.Book
}
