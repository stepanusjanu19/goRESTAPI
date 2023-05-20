package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stepanusjanu19/goRESTAPI/api/models"
	"github.com/stepanusjanu19/goRESTAPI/api/responses"
	"github.com/stepanusjanu19/goRESTAPI/api/utils/formaterror"
)

func (server *Server) CreateItem(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	item := models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	item.Prepare()
	err = item.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
}