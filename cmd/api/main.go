package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nelsonfrank/backend-api-go/internal/config"
	"github.com/nelsonfrank/backend-api-go/internal/repository"
	"github.com/nelsonfrank/backend-api-go/internal/services"
	transport "github.com/nelsonfrank/backend-api-go/internal/transport/http"
)

func main() {
	cfg := config.Load()

	conn, err := sql.Open("pgx", cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Initialize repositories
	repos := repository.NewRepositories(conn)

	// Initialize services
	services := services.NewServices(repos)

	// API
	api := transport.NewAPI(services)
	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", api.Router())
}
