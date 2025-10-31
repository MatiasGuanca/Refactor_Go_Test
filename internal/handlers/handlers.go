package handlers

import (
	//"context"
	"net/http"
	//"slices"

	//"educabot.com/bookshop/internal/models"

	"educabot.com/bookshop/internal/providers"
	services "educabot.com/bookshop/internal/service"
	"github.com/gin-gonic/gin"
)

type GetInformation struct {
	booksProvider    providers.BooksProvider
	booksInformation services.InformationService
}

func NewGetInformation(booksProvider providers.BooksProvider, booksInformation services.InformationService) GetInformation {
	return GetInformation{
		booksProvider,
		booksInformation,
	}
}

type GetRequest struct {
	Author string `form:"author" binding:"required"`
}

func (h GetInformation) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query GetRequest
		if err := ctx.ShouldBindQuery(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		reqCtx := ctx.Request.Context()
		books := h.booksProvider.GetBooks(reqCtx)

		result := h.booksInformation.BooksInformation(reqCtx, books, query.Author)
		if result["error"] != nil {
			ctx.JSON(http.StatusNotFound, result)
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
