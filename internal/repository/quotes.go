package repository

import (
	"database/sql"
	"quotetion_book/internal/models"
)

type QuotesRepository struct {
	db *sql.DB
}

func NewQuotesRepository(db *sql.DB) *QuotesRepository {
	return &QuotesRepository{db: db}
}
func (db *QuotesRepository) AddQuote(quote models.QuoteBook) error {
	_, err := db.db.Exec("INSERT INTO quote (author, quote) VALUES ($1, $2)",
		quote.Author, quote.Quote)

	if err != nil {
		return err
	}

	return nil
}

func (db *QuotesRepository) GetAllQuotes() (*sql.Rows, error) {
	rows, err := db.db.Query("SELECT id, author, quote FROM quote")

	if err != nil {
		return nil, err
	}

	return rows, nil
}
