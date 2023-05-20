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

func (server *Server) CreateItemGroup(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	itemgroup := models.ItemGroup{}
	err = json.Unmarshal(body, &itemgroup)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	itemgroup.Prepare()
	err = itemgroup.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	itemGroupCreated, err := itemgroup.SaveItemGroup(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, itemGroupCreated.ITEM_GROUP_ID))
	mapD := map[string]interface{}{"message": "Data Item Group Berhasil Ditambah", "data":itemGroupCreated}
	responses.JSON(w, http.StatusCreated, mapD)
}

func (server *Server) GetItemGroups(w http.ResponseWriter, r *http.Request)  {
	itemgroup := models.ItemGroup{}
	itemgroups, err := itemgroup.FindAllItemGroup(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, itemgroups)
}

func (server *Server) GetItemGroup(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	igid, err := strconv.ParseUint(vars["item_group_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	itemgroup := models.ItemGroup{}
	itemgroupGotten, err := itemgroup.FindItemGroupByID(server.DB, igid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, itemgroupGotten)
}

func (server *Server) UpdateItemGroup(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	igid, err := strconv.ParseUint(vars["item_group_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	
	itemgroup := models.ItemGroup{}
	err = server.DB.Debug().Model(models.ItemGroup{}).Where("item_group_id = ?", igid).Take(&itemgroup).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Item Group not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemgroupUpdate := models.ItemGroup{}
	err = json.Unmarshal(body, &itemgroupUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemgroupUpdate.Prepare()
	err = itemgroupUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemgroupUpdate.ITEM_GROUP_ID = itemgroup.ITEM_GROUP_ID
	itemgroupUpdated, err := itemgroupUpdate.UpdateItemGroup(server.DB, igid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	mapD := map[string]interface{}{"message": "Data Item Group Berhasil Diubah", "data":itemgroupUpdated}
	responses.JSON(w, http.StatusOK, mapD)
}

func (server *Server) DeleteItemGroup(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	igid, err := strconv.ParseUint(vars["item_group_id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	itemgroup := models.ItemGroup{}
	err = server.DB.Debug().Model(models.ItemGroup{}).Where("item_group_id = ?", igid).Take(&itemgroup).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Item Group Not Found"))
		return
	}

	_, err = itemgroup.DeleteItemGroup(server.DB, igid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", igid))
	mapD := map[string]string{"message": "Data Item Group Berhasil Dihapus"}
	responses.JSON(w, http.StatusOK, mapD)
}
