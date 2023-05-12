package controllers

import (
	"net/http"
	"github.com/stepanusjanu19/goRESTAPI/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request)  {
	responses.JSON(w, http.StatusOK, "Welcome REST CLIENT Go")
}