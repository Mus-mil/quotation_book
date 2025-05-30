package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"quotetion_book/internal/models"
)

func (h *Handlers) QuotesPost(ctx *gin.Context) {
	var quote models.QuoteBook

	if err := ctx.BindJSON(&quote); err != nil {
		log.Println("ошибка парсинга контекста")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ошибка": "неправильный формат данных",
		})
		return
	}

	if err := h.serv.AddQuote(quote); err != nil {
		log.Println("ошибка базы данных: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ошибка": "непревиденная ошибка базы данных",
		})
	}
}

func (h *Handlers) QuotesGet(ctx *gin.Context) {
	allQuotes, err := h.serv.GetAllQuotes()

	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ошибка": "непревиденная ошибка базы данных",
		})
		return
	}

	ctx.JSON(http.StatusOK, allQuotes)
}
