package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"user-service/helper"
	"user-service/middlewares"
	"user-service/models"

	"github.com/gorilla/mux"
)

func (server *Server) CreateUsers(w http.ResponseWriter, r *http.Request) {
	data := &models.StdUser{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = data.Validate("create")
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := data.SaveUsers(server.DB)
	if err != nil {
		format := helper.FormatError(err.Error())
		helper.ERROR(w, http.StatusUnprocessableEntity, format)
		return
	}

	userScope := models.StdUserScope{}
	_, err = userScope.SaveUserScope(server.DB, result.Uid, "user")
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, result.Uid))
	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusCreated,
		M{
			"data":    result,
			"message": "Successfully",
			"status":  http.StatusOK,
		},
	)
}

func (server *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	data := models.StdUser{}
	result, err := data.FindAllUsers(server.DB, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "data": result})
}

func (server *Server) GetUsersByID(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	uid := vars["id"]

	data := models.StdUser{}
	result, err := data.FindUsersByID(server.DB, uid, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "data": result})
}

func (server *Server) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	uid := vars["id"]

	data := &models.StdUser{}
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = data.SaveUpdateUsers(server.DB, uid, token)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "message": "Successfully"})
}

func (server *Server) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	uid := vars["id"]

	data := models.StdUser{}
	_, err = data.SaveDeleteUsers(server.DB, uid, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "message": "Successfully"})
}
