package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"educabot.com/bookshop/internal/providers"
	services "educabot.com/bookshop/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMetrics_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	booksProvider := providers.NewHttpBooksProvider("https://6781684b85151f714b0aa5db.mockapi.io/api/v1")
	metricsService := services.NewInformationService()

	handler := NewGetInformation(booksProvider, metricsService)

	r := gin.Default()
	r.GET("/", handler.Handle())

	req := httptest.NewRequest(http.MethodGet, "/?author=Alan+Donovan", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var resBody map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &resBody)

	assert.Equal(t, 11000, int(resBody["mean_units_sold"].(float64)))
	assert.Equal(t, "The Go Programming Language", resBody["cheapest_book"])
	assert.Equal(t, 1, int(resBody["books_written_by_author"].(float64)))
}
