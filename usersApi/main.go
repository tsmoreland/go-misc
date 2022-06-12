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

func migrateDb(filename string) error {
	_ = os.Remove(filename)
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}

	var userRepository shared.Repository = infrastructure.NewSQLiteRepository(db)
	defer func() {
		err = userRepository.Close()
	}()

	err = userRepository.Migrate()
	if err != nil {
		return err
	}

	users, err := userRepository.All()
	if err != nil {
		for _, user := range users {
			fmt.Printf("%d = %s\n", user.ID, user.Name)
		}
	}

	return nil
}

func main() {

	if Version != "" {
		fmt.Printf("Version: %s", Version)
	}

	config := app.
		NewConfigurationBuilder().
		AddJsonFile("configuration.json").
		AddEnvironment().
		Build()

	err := migrateDb(config.DatabaseFile())
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
		HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = w.Write([]byte("{}"))
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
