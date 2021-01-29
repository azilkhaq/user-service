package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"user-service/helper"
	"user-service/middlewares"
	"user-service/models"

	"github.com/gorilla/mux"
)

func (server *Server) GetAllUserScope(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	data := models.StdUserScope{}
	result, err := data.FindAllUserScope(server.DB, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "data": result})
}

func (server *Server) GetUserScopeByID(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := models.StdUserScope{}
	result, err := data.FindUserScopeByID(server.DB, ID, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "data": result})
}

func (server *Server) UpdateUserScope(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := &models.StdUserScope{}
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = data.SaveUpdateUserScope(server.DB, ID, token)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "message": "Successfully"})
}
