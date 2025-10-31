package services_test

import (
	"context"
	"testing"

	"educabot.com/bookshop/internal/models"
	"educabot.com/bookshop/internal/service"
)

func TestBooksInformation_NoBooks(t *testing.T) {
	svc := services.NewInformationService()
	got := svc.BooksInformation(context.Background(), []models.Book{}, "any")

	if got["error"] != "no books available" {
		t.Errorf("expected error message, got %+v", got)
	}
}

func TestBooksInformation_WithBooks(t *testing.T) {
	svc := services.NewInformationService()
	books := []models.Book{
		{Name: "Book A", Author: "John", Price: 10, UnitsSold: 100},
		{Name: "Book B", Author: "Jane", Price: 5, UnitsSold: 200},
		{Name: "Book C", Author: "John", Price: 15, UnitsSold: 50},
	}

	got := svc.BooksInformation(context.Background(), books, "John")

	if got["mean_units_sold"] != uint(116) { // (100+200+50)/3
		t.Errorf("expected mean 116, got %v", got["mean_units_sold"])
	}

	if got["cheapest_book"] != "Book B" {
		t.Errorf("expected cheapest 'Book B', got %v", got["cheapest_book"])
	}

	if got["books_written_by_author"] != uint(2) {
		t.Errorf("expected 2 books by John, got %v", got["books_written_by_author"])
	}
}
