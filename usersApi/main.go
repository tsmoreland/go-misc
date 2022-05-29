package main

import (
	"usersApi/domain"
	"usersApi/infrastructure"

	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbFilename = "users.db"

func main() {
	_ = os.Remove(dbFilename)
	db, err := sql.Open("sqlite3", dbFilename)
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
