package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	//"usersApi/domain"
)

type UsersHandler struct {
	version string
}

func NewUsersHandler(version string) *UsersHandler {
	h := UsersHandler{version: version}
	return &h
}

func (h UsersHandler) ConfigureHandlers(router *mux.Router) {

	router.HandleFunc("/api/v1/users", h.HandlerUsersRoot)
	router.HandleFunc("/api/v1/users/{id}", h.HandlerUsersRoot)
}

func (h UsersHandler) HandlerUsersRoot(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
	}
}

func (h UsersHandler) HandlerUsersId(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write([]byte("{\"error\":\"invalid id, expected a number\"}"))
		log.Print(err.Error())
		return
	}

	_ = id

	switch req.Method {
	case http.MethodGet:
	case http.MethodPut:
	case http.MethodDelete:
	}
}
