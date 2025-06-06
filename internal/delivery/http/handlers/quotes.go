package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"quotetion_book/internal/models"
	"strconv"
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
	if quote.Author == "" || quote.Quote == "" {
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

	ctx.Status(http.StatusCreated)
}

func (h *Handlers) QuotesGet(ctx *gin.Context) {
	author := ctx.Query("author")

	if author != "" {
		QuotesAuthor, err := h.serv.GetQuotesFromAuthor(author)

		if QuotesAuthor == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ошибка": "автора нет в базе данных",
			})
			return
		}

		if err != nil {
			log.Println(err)

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"ошибка": "непревиденная ошибка базы данных",
			})
			return
		}

		ctx.JSON(http.StatusOK, QuotesAuthor)
	} else {
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

}

func (h *Handlers) RandomQuote(ctx *gin.Context) {
	randomQuote, err := h.serv.GetRandomQuote()

	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ошибка": "непревиденная ошибка базы данных",
		})
		return
	}

	ctx.JSON(http.StatusOK, randomQuote)
}

func (h *Handlers) QuotesDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"ошибка": "введите число",
		})
		return
	}

	err = h.serv.DeleteQuotesFromID(id)

	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ошибка": "id не существует",
		})
		return
	}
}
