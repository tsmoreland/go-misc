package main

import (
	"fmt"
	"usersApi/app"
	"usersApi/domain"
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

	config := app.NewConfigurationBuilder().AddJsonFile("configuration.json").AddEnvironment().Build()

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

	userA := domain.User{
		Name:     "A",
		FullName: "Alpha",
	}

	userB := domain.User{
		Name:     "B",
		FullName: "Bravo",
	}

	createdA, err := userRepository.Create(userA)
	if err != nil {
		log.Fatal(err)
	}

	createdB, err := userRepository.Create(userB)
	if err != nil {
		log.Fatal(err)
	}

	err = userRepository.Delete(createdA.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = userRepository.Delete(createdB.ID)
	if err != nil {
		log.Fatal(err)
	}

	/*
		mux := http.NewServeMux()
		http.ListenAndServe(fmt.Sprint(":%d", config.Port()), mux)
	*/
}

/*

func main() {
	err = userRepository.Delete(createdA.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = userRepository.Delete(createdB.ID)
	if err != nil {
		log.Fatal(err)
	}
}


*/
