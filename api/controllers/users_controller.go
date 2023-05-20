package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "github.com/stepanusjanu19/goRESTAPI/api/auth"
	"github.com/stepanusjanu19/goRESTAPI/api/models"
	"github.com/stepanusjanu19/goRESTAPI/api/responses"
	"github.com/stepanusjanu19/goRESTAPI/api/utils/formaterror"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.USER_ID))
	mapD := map[string]interface{}{"message": "Data Users Berhasil Ditambah", "data":userCreated}
	responses.JSON(w, http.StatusCreated, mapD)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindByUserID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("user_id = ?", uid).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	userUpdate:= models.User{}
	err = json.Unmarshal(body, &userUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// tokenID, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if tokenID != uint32(uid) {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }

	userUpdate.Prepare()
	err = userUpdate.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userUpdate.USER_ID = user.USER_ID
	updatedUser, err := userUpdate.UpdateUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	// responses.JSON(w, http.StatusOK, updatedUser)
	mapD := map[string]interface{}{"message": "Data Users Berhasil Diubah", "data":updatedUser}
	responses.JSON(w, http.StatusOK, mapD)

}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// tokenID, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if tokenID != 0 && tokenID != uint32(uid) {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }

	err = server.DB.Debug().Model(models.User{}).Where("user_id = ?", uid).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User Not Found"))
		return
	}

	_, err = user.DeleteUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	mapD := map[string]string{"message": "Data Users Berhasil Dihapus"}
	responses.JSON(w, http.StatusOK, mapD)
}
