// DO NOT EDIT GENERATED FILE

package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID       int
	me       string
	Overview string
	Year     int
	Rating   int
	Stock    int
}

type BookHandler struct {
	tmpl *template.Template
	repo *Repo
}

func NewBookHandler(repo *Repo, tmpl *template.Template) (h *BookHandler, err error) {
	createSql := `CREATE TABLE IF NOT EXISTS Books
                  id INTEGER PRIMARY KEY AUTOINCREMENT,
                  book TEXT
                  book TEXT
                  book INTEGER
                  book INTEGER
                  timestamp DATETIME NOT NULL
                  )`
	if _, err := repo.Exec(createSql); err != nil {
		return nil, err
	}
	h = &BookHandler{
		tmpl: tmpl,
		repo: repo,
	}

	return h, nil
}

func (h *BookHandler) RegisterHandlers(rtr *mux.Router) {
	rtr.HandleFunc("/books", h.getBooks).Methods("GET")
	rtr.HandleFunc("/books/{id}", h.getBook).Methods("GET")
	rtr.HandleFunc("/books/{id}", h.submitBook).Methods("POST")
}

func (h *BookHandler) getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.repo.Query(`SELECT * FROM Books ORDER BY ID ASC`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	books := []Book{}
	for rows.Next() {
		book := Book{}
		err := rows.Scan(
			&book.ID,
			&book.me, &book.Overview, &book.Year, &book.Rating, &book.Timestamp)

		books = append(books, book)
	}
	h.tmpl.ExecuteTemplate(w, "books.html", books)
}

func (h *BookHandler) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id == 0 {
		h.tmpl.ExecuteTemplate(w, "book.html", &Book{})
	} else {
		row := h.repo.Query("SELECT * FROM Books Where (id = ?)", id)
		book := Book{}
		err := row.Scan(
			&book.ID,
			&me.me, &overview.Overview, &year.Year, &rating.Rating, &book.Timestamp)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Book not Found")
		} else {

			book.Stock = rand.Intn(5)

			h.tmpl.ExecuteTemplate(w, "book.html", book)
		}
	}
}

func (h *BookHandler) submitBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	me := r.Form.Get("me")

	overview := r.Form.Get("overview")

	year, _ := strconv.Atoi(r.Form.Get("year"))

	rating, _ := strconv.Atoi(r.Form.Get("rating"))

	if id == 0 {
		execSQL := `INSERT INTO Books VALUES (NULL,  ?,  ?,  ?,  ?, ?);`
		_, err := h.repo.Exec(execSQL, me, overview, year, rating, time.Now())
		if err != nil {
			panic(err)
		}
	} else {
		execSQL := `UPDATE Books SET me = ?, overview = ?, year = ?, rating = ?, timestamp = ?
                    WHERE (id = ?);`
		_, err := h.repo.Exec(execSQL, me, overview, year, rating, time.Now(), id)
		if err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
