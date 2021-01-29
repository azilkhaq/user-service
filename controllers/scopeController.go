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

func (server *Server) CreateScope(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	data := &models.StdScope{}
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := data.SaveScope(server.DB, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	userScope := models.StdUserScope{}
	_, err = userScope.SaveUserScope(server.DB, result.Scope, "scope")
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, result.ID))
	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusCreated,
		M{
			"data":    result,
			"message": "Successfully",
			"status":  http.StatusOK,
		},
	)
}

func (server *Server) GetAllScope(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	data := models.StdScope{}
	result, err := data.FindAllScope(server.DB, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "data": result})
}

func (server *Server) DeleteScope(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := models.StdScope{}
	_, err = data.SaveDeleteScope(server.DB, ID, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "message": "Successfully"})
}
