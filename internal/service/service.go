package service

import (
	"quotetion_book/internal/models"
	"quotetion_book/internal/repository"
)

type QuotationService interface {
	AddQuote(quote models.QuoteBook) error
	GetAllQuotes() ([]models.QuoteBookID, error)
	GetRandomQuote() (models.QuoteBook, error)
	GetQuotesFromAuthor(author string) ([]models.QuoteBookID, error)
	DeleteQuotesFromID(id int) error
}

type Service struct {
	QuotationService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		QuotationService: NewQuoteService(repo.Quotation),
	}
}
