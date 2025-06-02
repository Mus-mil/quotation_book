package handlers

import (
	"github.com/gin-gonic/gin"
	"quotetion_book/internal/service"
)

type Handlers struct {
	serv *service.Service
}

func NewHandlers(serv *service.Service) *Handlers {
	return &Handlers{serv: serv}
}

func RegisterRoutes(h *Handlers) *gin.Engine {
	router := gin.Default()

	router.POST("/quotes", h.QuotesPost)
	router.GET("/quotes", h.QuotesGet)
	router.GET("/quotes/random", h.RandomQuote)
	return router
}
