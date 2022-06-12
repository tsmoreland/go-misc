package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	"usersApi/app"
	"usersApi/infrastructure"
	"usersApi/shared"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Version string
)

func main() {

	if Version != "" {
		fmt.Printf("Version: %s", Version)
	}

	config := app.
		NewConfigurationBuilder().
		AddJsonFile("configuration.json").
		AddEnvironment().
		Build()

	_ = os.Remove(config.DatabaseFile())
	db, err := sql.Open("sqlite3", config.DatabaseFile())
	if err != nil {
		log.Fatal(err)
	}

	var userRepository shared.Repository = infrastructure.NewSQLiteRepository(db)
	defer userRepository.Close()

	err = userRepository.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	//r.Schemes("http")

	var repositoryFactory shared.RepositoryFactory = infrastructure.
		NewSqliteRepositoryFactory(config.DatabaseFile())

	userHandler := app.NewUsersHandler(Version, repositoryFactory)
	userHandler.ConfigureHandlers(r)

	r.
		HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("Hello"))
		}).
		Methods("GET")
	r.
		HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write([]byte("{}"))
		}).
		Methods("GET")

	serverAddress := fmt.Sprintf(":%d", config.Port())
	server := &http.Server{
		Handler:      r,
		Addr:         serverAddress,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	//http.ListenAndServe(serverAddress, r)

	log.Fatal(server.ListenAndServe())
}
