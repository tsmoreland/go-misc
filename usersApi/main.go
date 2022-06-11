package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"usersApi/app"
	"usersApi/infrastructure"

	"database/sql"
	"log"
	"os"

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

	userRepository := infrastructure.NewSQLiteRepository(db)

	err = userRepository.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.
		Methods("GET", "POST", "PUT", "DELETE").
		Schemes("http")

	userHandler := app.NewUsersHandler(Version)
	userHandler.ConfigureHandlers(r)

	serverAddress := fmt.Sprint(":%d", config.Port())
	server := &http.Server{
		Handler:      r,
		Addr:         serverAddress,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
