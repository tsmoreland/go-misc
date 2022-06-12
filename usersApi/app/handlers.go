package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"usersApi/shared"
	//"usersApi/domain"
)

type UsersHandler struct {
	version           string
	repositoryFactory shared.RepositoryFactory
}

func NewUsersHandler(version string, repositoryFactory shared.RepositoryFactory) *UsersHandler {
	h := UsersHandler{version: version, repositoryFactory: repositoryFactory}
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

func validateId(idParameter string, res http.ResponseWriter) (int, bool) {
	id, err := strconv.Atoi(idParameter)
	if err == nil {
		return id, true
	}

	log.Print(err.Error())
	res.WriteHeader(http.StatusBadRequest)
	_, err = res.Write([]byte("{\"error\":\"invalid id, expected a number\"}"))
	log.Print(err.Error())
	return 0, false
}

func internalServerError(res http.ResponseWriter, message string) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", message)))

}

func failed(statusCode int, res http.ResponseWriter, message string) {
	res.WriteHeader(statusCode)
	res.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", message)))
}
func successful(res http.ResponseWriter, statusCode int, content []byte) {
	res.WriteHeader(statusCode)
	res.Write(content)
}

func (h UsersHandler) HandlerUsersId(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(req)

	id, success := validateId(vars["id"], res)
	if !success {
		return
	}

	r, err := h.repositoryFactory.Build()
	if err != nil {
		log.Print(err.Error())
		internalServerError(res, "unable to open connection to the database")
		return
	}

	_ = id
	_ = r

	switch req.Method {
	case http.MethodGet:
		u, err := r.GetById(int64(id))
		if err != nil {
			failed(http.StatusNotFound, res, "user not found")
		} else {
			serialized, err := json.Marshal(u)
			// TODO: rework this to pass the u in as an any and let succeeded handle the serialize failure
			if err != nil {
				log.Print(err.Error())
				internalServerError(res, "unable to serialized user")
			} else {
				successful(res, http.StatusOK, serialized)
			}
		}

	case http.MethodPut:
	case http.MethodDelete:
	}
}
