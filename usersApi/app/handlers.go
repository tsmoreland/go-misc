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
	router.HandleFunc("/api/v1/users/{id}", h.HandlerUsersId)
}

func (h UsersHandler) HandlerUsersRoot(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	r, err := h.repositoryFactory.Build()
	if err != nil {
		writeError(http.StatusInternalServerError, err, res, "unable to open connection to the database")
		return
	}
	defer func() {
		err = r.Close()
	}()

	switch req.Method {
	case http.MethodGet:
		users, err := r.All()
		if err != nil {
			writeError(http.StatusInternalServerError, err, res, "unable to retrieve users")
		}

		var dtos []UserSummaryDto
		for _, user := range users {
			dtos = append(dtos, *NewUserSummaryDto(user.ID, user.Name))
		}
		writeResponse(res, http.StatusOK, dtos)

	case http.MethodPost:
	}
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
		writeError(http.StatusInternalServerError, err, res, "unable to open connection to the database")
		return
	}
	defer func() {
		err = r.Close()
	}()

	u, err := r.GetById(int64(id))
	if err != nil {
		writeError(http.StatusNotFound, err, res, "user not found")
	}

	switch req.Method {
	case http.MethodGet:
		writeResponse(res, http.StatusOK, NewUserDto(*u))
	case http.MethodPut:
		// TODO: we need to update u with a body of request

	case http.MethodDelete:
		_ = r.Delete(int64(id))
		writeResponse(res, http.StatusNoContent, nil)
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

func writeError(statusCode int, err error, res http.ResponseWriter, message string) {
	if err != nil {
		log.Print(err.Error())
	}

	res.WriteHeader(statusCode)
	_, _ = res.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", message)))
}
func writeResponse(res http.ResponseWriter, statusCode int, content any) {
	if content == nil {
		res.WriteHeader(statusCode)
		return
	}

	serialized, err := json.Marshal(content)
	if err != nil {
		writeError(http.StatusInternalServerError, err, res, "serialization failure")
	} else {
		res.WriteHeader(statusCode)
		_, _ = res.Write(serialized)
	}
}
