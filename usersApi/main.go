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
	// TODO:
	// add method, maybe to factory handle the db migration and update repository to include a Close that can be
	// deferred
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var userRepository shared.Repository = infrastructure.NewSQLiteRepository(db)

	err = userRepository.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.
		Methods("GET", "POST", "PUT", "DELETE").
		Schemes("http")

	var repositoryFactory shared.RepositoryFactory = infrastructure.
		NewSqliteRepositoryFactory(config.DatabaseFile())

	userHandler := app.NewUsersHandler(Version, repositoryFactory)
	userHandler.ConfigureHandlers(r)

	serverAddress := fmt.Sprintf(":%d", config.Port())
	server := &http.Server{
		Handler:      r,
		Addr:         serverAddress,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
