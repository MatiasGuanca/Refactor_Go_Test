package services

import (
	"context"
	"slices"

	"educabot.com/bookshop/internal/models"
)

type InformationService struct{}

func NewInformationService() InformationService {
	return InformationService{}
}

func (s InformationService) BooksInformation(ctx context.Context, books []models.Book, author string) map[string]interface{} {
	if len(books) == 0 {
		return map[string]interface{}{
			"error": "no books available",
		}
	}

	return map[string]interface{}{
		"mean_units_sold":         s.meanUnitsSold(ctx, books),
		"cheapest_book":           s.cheapestBook(ctx, books).Name,
		"books_written_by_author": s.booksWrittenByAuthor(ctx, books, author),
	}
}

func (s InformationService) meanUnitsSold(_ context.Context, books []models.Book) uint {
	var sum uint
	for _, b := range books {
		sum += b.UnitsSold
	}
	return sum / uint(len(books))
}

func (s InformationService) cheapestBook(_ context.Context, books []models.Book) models.Book {
	return slices.MinFunc(books, func(a, b models.Book) int {
		return int(a.Price - b.Price)
	})
}

func (s InformationService) booksWrittenByAuthor(_ context.Context, books []models.Book, author string) uint {
	var count uint
	for _, b := range books {
		if b.Author == author {
			count++
		}
	}
	return count
}
