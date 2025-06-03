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

func (db *QuotesRepository) GetRowsCount() (int, error) {
	rowsCount := 0
	err := db.db.QueryRow("SELECT COUNT(*) FROM quote").Scan(&rowsCount)

	if err != nil {
		return 0, err
	}

	return rowsCount, nil
}

func (db *QuotesRepository) GetQuoteFromID(offset int) (models.QuoteBook, error) {
	var q models.QuoteBook

	err := db.db.QueryRow("SELECT author, quote  FROM quote LIMIT 1 OFFSET $1", offset).Scan(&q.Author, &q.Quote)

	if err != nil {
		return models.QuoteBook{}, err
	}

	return q, nil
}

func (db *QuotesRepository) GetQuotesFromAuthor(author string) (*sql.Rows, error) {
	rows, err := db.db.Query("SELECT id, author, quote FROM quote WHERE author = $1", author)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (db *QuotesRepository) DeleteQuoteFromID(id int) error {
	_, err := db.db.Exec("DELETE FROM quote WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
