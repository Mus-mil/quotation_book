package service

import (
	"quotetion_book/internal/models"
	"quotetion_book/internal/repository"
)

type Quotation interface {
	AddQuote(quote models.QuoteBook) error
	GetAllQuotes() ([]models.QuoteBookID, error)
	GetRandomQuote() (models.QuoteBook, error)
	GetQuotesFromAuthor(author string) ([]models.QuoteBookID, error)
}

type Service struct {
	Quotation
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Quotation: NewQuoteService(repo.Quotation),
	}
}
