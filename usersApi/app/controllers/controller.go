package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	Configure(r *mux.Router) Controller
}

func ok(w http.ResponseWriter, content any) {
	writeResponse(w, http.StatusOK, content)
}

func created(w http.ResponseWriter, content any) {
	writeResponse(w, http.StatusCreated, content)
}

func writeResponse(w http.ResponseWriter, statusCode int, content any) {
	if content == nil {
		w.WriteHeader(statusCode)
		return
	}

	serialized, err := json.Marshal(content)
	if err != nil {
		// TODO: problem details
		//writeError(http.StatusInternalServerError, err, res, "serialization failure")
	} else {
		w.WriteHeader(statusCode)
		_, _ = w.Write(serialized)
	}
}
