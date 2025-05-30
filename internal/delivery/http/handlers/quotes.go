package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"quotetion_book/internal/models"
)

func (h *Handlers) QuotesPost(ctx *gin.Context) {
	var quote models.QuoteBook

	if err := ctx.BindJSON(&quote); err != nil {
		log.Println("ошибка парсинга контекста")
		return
	}

	if err := h.serv.AddQuote(quote); err != nil {
		log.Println("ошибка базы данных: ", err)
	}
}
