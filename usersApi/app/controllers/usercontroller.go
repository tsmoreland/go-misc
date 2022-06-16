package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"usersApi/app/middleware"
	"usersApi/app/models"
	"usersApi/shared"
)

type userController struct {
	factory shared.RepositoryFactory
}

func NewUserController(factory shared.RepositoryFactory) Controller {
	return &userController{factory: factory}
}

func (c *userController) Configure(r *mux.Router) Controller {

	sr := r.PathPrefix("/users").Subrouter()
	sr.HandleFunc("/", c.root)

	sr.Use(middleware.SetJsonContentMiddleware)

	return c
}

func (c userController) root(w http.ResponseWriter, r *http.Request) {

	repository, err := c.factory.Build()
	if err != nil {
		// TODO add problemdetails to model with a method to write error response with given response code
		return
	}

	switch r.Method {
	case http.MethodGet:
		c.getRoot(w, r, repository)
	case http.MethodPost:
		c.postRoot(w, r, repository)
	}

}

func (c userController) getRoot(w http.ResponseWriter, r *http.Request, repository shared.Repository) {

	users, err := repository.All()
	if err != nil {
		// TODO add problemdetails to model with a method to write error response with given response code
		// looking for something like NewProblemDetails(500, title, description).Write(w)
		return
	}

	var dtos []models.UserSummaryDto
	for _, user := range users {
		dtos = append(dtos, *models.NewUserSummaryDto(user.ID, user.Name))
	}
	ok(w, dtos)
}
func (c userController) postRoot(w http.ResponseWriter, r *http.Request, repository shared.Repository) {
	var dto models.UserInputDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		// TODO: problemd details
		//writeError(http.StatusBadRequest, err, res, "input input")
	}

	user, err := repository.Create(*dto.BuildUser())
	if err != nil {
		// TODO:  problem details
		//writeError(http.StatusInternalServerError, err, res, "failed to add user")
	} else {
		created(w, models.NewUserDto(*user))
	}
}
