package infrastructure

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"usersApi/domain"
)

var (
	ErrorDoesNotExist = errors.New("item not found")
	ErrorUpdateFailed = errors.New("update failed")
	ErrorDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (repository *SQLiteRepository) Migrate() error {
	command := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		fullName TEXT NOT NULL
	);
	`
	_, err := repository.db.Exec(command)
	return err
}

func (repository *SQLiteRepository) Create(user domain.User) (*domain.User, error) {

	command := `
	INSERT INTO users (name, fullName)
	values (?, ?)
	`

	result, err := repository.db.Exec(command, user.Name, user.FullName)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = id

	return &user, nil
}

func (repository *SQLiteRepository) All() ([]domain.User, error) {
	rows, err := repository.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer func() {
		if tempErr := rows.Close(); tempErr != nil {
			err = tempErr
		}
	}()

	var all []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.FullName); err != nil {
			return nil, err
		}
		all = append(all, user)
	}
	return all, nil
}

func (repository *SQLiteRepository) GetById(id int64) (*domain.User, error) {
	row := repository.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.FullName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorDoesNotExist
		}
		return nil, err
	}
	return &user, nil
}

func (repository *SQLiteRepository) GetByName(name string) (*domain.User, error) {
	row := repository.db.QueryRow("SELECT * FROM users WHERE name = ?", name)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.FullName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorDoesNotExist
		}
		return nil, err
	}
	return &user, nil
}

func (repository *SQLiteRepository) Update(id int64, updated domain.User) (*domain.User, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}

	command := `
	UPDATE users 
	SET name = ?, fullName = ?
	WHERE id = ?
	`

	res, err := repository.db.Exec(command, updated.Name, updated.FullName, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrorUpdateFailed
	}

	return &updated, nil
}

func (repository *SQLiteRepository) Delete(id int64) error {
	res, err := repository.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorDeleteFailed
	}

	return err
}
