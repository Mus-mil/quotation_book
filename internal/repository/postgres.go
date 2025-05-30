package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"quotetion_book/internal/configs"
)

func InitDB(cfg configs.PostgresConfig) *sql.DB {
	connectDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name, cfg.Sslmode)
	log.Println("connect db:", connectDB)
	db, err := sql.Open("postgres", connectDB)
	if err != nil {
		log.Fatal("not authorize postgres")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке соединения: %v", err)
	}

	return db
}
