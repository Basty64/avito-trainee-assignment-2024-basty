package app

import (
	"avito-trainee-assignment-2024-basty/internal/handlers"
	"avito-trainee-assignment-2024-basty/internal/repository/postgres"
	"avito-trainee-assignment-2024-basty/internal/services"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

type App struct {
	err              error
	router           http.Handler
	databaseURL      string
	handlers         *handlers.Handlers
	service          *services.Service
	bannerRepository *postgres.BannersRepository
}

func NewApp() (*App, error) {
	a := &App{}

	a.databaseURL = "postgres://postgres:postgres@localhost:5433/postgres"

	postgresPool, err := pgxpool.New(context.Background(), a.databaseURL)
	if err != nil {
		panic(err)
	}

	if err := postgresPool.Ping(context.Background()); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}

	defer postgresPool.Close()

	a.bannerRepository = postgres.NewBannersRepository(postgresPool)

	a.databaseURL = "postgres://postgres:postgres@localhost:5432/postgres"

	a.service = services.NewService(a.bannerRepository)
	return a, nil

}

func (a *App) HTTPRouter() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/banner", handlers.POSTBannerHandler(a.service)).Methods(http.MethodPost)

	//router.Handle("/banner", handler).Methods("GET")
	//router.Handle("/user_banner", handler).Methods("GET")
	//router.Handle("/banner/{id}", handler).Methods("PATCH")
	//router.Handle("/banner/{id}", handler).Methods("DELETE")

	a.router = router
	return a.router
}

func (a *App) Run() error {

	port := "localhost:8080"
	fmt.Println("Server is running on port " + port)

	server := &http.Server{Addr: port, Handler: a.HTTPRouter()}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
