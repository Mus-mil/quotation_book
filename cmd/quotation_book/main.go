package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"quotetion_book/internal/configs"
	"quotetion_book/internal/delivery/http/handlers"
	"quotetion_book/internal/repository"
	"quotetion_book/internal/server"
	"quotetion_book/internal/service"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	if err := initConfigs(); err != nil {
		log.Print("No yaml file found")
	}
	cfg := configs.NewConfigs()

	db := repository.InitDB(cfg.Postgres)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Not closer DB")
		}
	}(db)

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handler := handlers.NewHandlers(serv)

	router := handlers.RegisterRoutes(handler)

	srv := server.Server{}
	func() {
		err := srv.RunServer("8080", router)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func initConfigs() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
