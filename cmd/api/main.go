package main

import (
	"fmt"

	"educabot.com/bookshop/internal/handlers"
	"educabot.com/bookshop/internal/providers"
	services "educabot.com/bookshop/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.SetTrustedProxies(nil)

	booksProvider := providers.NewHttpBooksProvider("https://6781684b85151f714b0aa5db.mockapi.io/api/v1")
	metricsService := services.NewInformationService()

	handler := handlers.NewGetInformation(booksProvider, metricsService)
	router.GET("/metrics", handler.Handle())
	router.Run(":3000")
	fmt.Println("Starting server on :3000")
}
