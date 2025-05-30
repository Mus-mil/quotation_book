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

func (r *QuoteService) GetAllQuotes() ([]models.QuoteBookID, error) {
	rows, err := r.repo.GetAllQuotes()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.QuoteBookID

	for rows.Next() {
		var q models.QuoteBookID
		if err := rows.Scan(&q.ID, &q.Author, &q.Quote); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	return quotes, nil
}
