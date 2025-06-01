package repository

import (
	"database/sql"
	"quotetion_book/internal/models"
)

type Quotation interface {
	AddQuote(quote models.QuoteBook) error
	GetAllQuotes() (*sql.Rows, error)
	GetQuoteFromID(offset int) (models.QuoteBook, error)
	GetRowsCount() (int, error)
}

type Repository struct {
	Quotation
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Quotation: NewQuotesRepository(db),
	}
}
