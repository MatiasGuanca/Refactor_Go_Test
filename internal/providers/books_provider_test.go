package providers_test

import (
	"context"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"educabot.com/bookshop/internal/models"
	"educabot.com/bookshop/internal/providers"
)

func TestGetBooks_Success(t *testing.T) {
	expected := []models.Book{{Name: "Test Book", Author: "Matias", Price: 10, UnitsSold: 5}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(expected)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}))
	defer server.Close()

	p := providers.NewHttpBooksProvider(server.URL)
	got := p.GetBooks(context.Background())

	if len(got) != 1 || got[0].Name != "Test Book" {
		t.Errorf("expected one book, got %+v", got)
	}
}

func TestGetBooks_Non200Status(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	p := providers.NewHttpBooksProvider(server.URL)
	got := p.GetBooks(context.Background())
	if len(got) != 0 {
		t.Errorf("expected no books, got %+v", got)
	}
}

func TestGetBooks_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	p := providers.NewHttpBooksProvider(server.URL)
	got := p.GetBooks(context.Background())
	if len(got) != 0 {
		t.Errorf("expected no books, got %+v", got)
	}
}

func TestGetBooks_ReadError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		conn, _, _ := w.(http.Hijacker).Hijack() // forza cierre de conexión abrupta
		conn.Close()
	}))
	defer server.Close()

	p := providers.NewHttpBooksProvider(server.URL)
	got := p.GetBooks(context.Background())
	if len(got) != 0 {
		t.Errorf("expected empty result on read error, got %+v", got)
	}
}
