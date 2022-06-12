package infrastructure

import (
	"database/sql"
	"log"
	"usersApi/shared"
)

type SqliteRepositoryFactory struct {
	databaseFile string
}

func NewSqliteRepositoryFactory(databaseFile string) *SqliteRepositoryFactory {
	// TODO: ensure file exists

	return &SqliteRepositoryFactory{databaseFile: databaseFile}
}

func (f SqliteRepositoryFactory) Build() (shared.Repository, error) {

	db, err := sql.Open("sqlite3", f.databaseFile)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	var r shared.Repository = NewSQLiteRepository(db)
	return r, nil
}
