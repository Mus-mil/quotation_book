package service

import (
	"quotetion_book/internal/models"
	"quotetion_book/internal/repository"
)

type QuoteService struct {
	repo repository.Quotation
}

func NewQuoteService(repo repository.Quotation) *QuoteService {
	return &QuoteService{repo: repo}
}

func (r *QuoteService) AddQuote(quote models.QuoteBook) error {
	return r.repo.AddQuote(quote)
}
