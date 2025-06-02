package service

import (
	"math/rand/v2"
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
	defer rows.Close()

	if err != nil {
		return nil, err
	}

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

func (r *QuoteService) GetRandomQuote() (models.QuoteBook, error) {
	rowsCount, err := r.repo.GetRowsCount()

	if err != nil {
		return models.QuoteBook{}, err
	}

	offset := rand.IntN(rowsCount)

	var q models.QuoteBook
	q, err = r.repo.GetQuoteFromID(offset)

	if err != nil {
		return models.QuoteBook{}, err
	}

	return q, nil
}

func (r *QuoteService) GetQuotesFromAuthor(author string) ([]models.QuoteBookID, error) {
	rows, err := r.repo.GetQuotesFromAuthor(author)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

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
